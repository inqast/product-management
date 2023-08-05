package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"route256/libs/broker/kafka"
	dbmanager "route256/libs/db/manager"
	log "route256/libs/logger"
	"route256/libs/mw/logging"
	"route256/libs/mw/metrics"
	"route256/libs/mw/recovering"
	"route256/libs/mw/tracing"
	"route256/libs/mw/validation"
	"route256/libs/tracer"
	"route256/loms/internal/config"
	"route256/loms/internal/domain/service"
	orderRepo "route256/loms/internal/repository/order"
	stocksRepo "route256/loms/internal/repository/stocks"
	"route256/loms/internal/sender"
	"route256/loms/internal/server"
	api "route256/loms/pkg/loms/v1"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to init config: ", err)
	}

	if err := tracer.InitGlobal(config.AppConfig.Name); err != nil {
		log.Fatal("failed to init tracer: ", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.AppConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go runRest(config.AppConfig, lis)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := pgxpool.Connect(ctx, config.AppConfig.DBConString)
	if err != nil {
		log.Fatalf("connect to db: %v", err)
	}

	db := dbmanager.New(pool)

	stocksRepository := stocksRepo.New(db)
	orderRepository := orderRepo.New(db)

	producer, err := kafka.NewProducer(config.AppConfig.Kafka.Brokers)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}

	sender := sender.New(
		producer,
		config.AppConfig.Kafka.Topic,
	)

	lomsService := service.New(
		stocksRepository,
		orderRepository,
		db,
		sender,
	)

	lomsServer := server.New(lomsService)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovering.Interceptor,
			logging.Interceptor,
			validation.Interceptor,
			tracing.Interceptor,
			metrics.Interceptor,
		),
	)
	api.RegisterLomsServer(grpcServer, lomsServer)

	log.Infof("Serving gRPC on %s\n", lis.Addr().String())
	log.Fatal(grpcServer.Serve(lis))
}

func runRest(cfg config.Config, lis net.Listener) {
	conn, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "x-trace-id":
				return key, true
			}
			return runtime.DefaultHeaderMatcher(key)
		}),
	)

	if err := mux.HandlePath(http.MethodGet, "/metrics", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	}); err != nil {
		log.Fatal("something wrong with metrics handler", err)
	}

	err = api.RegisterLomsHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HttpPort),
		Handler: mux,
	}

	log.Infof("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	log.Fatal(gwServer.ListenAndServe())
}
