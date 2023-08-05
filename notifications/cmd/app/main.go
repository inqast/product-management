package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"route256/libs/broker"
	"route256/libs/broker/kafka"
	"route256/libs/cache"
	dbmanager "route256/libs/db/manager"
	log "route256/libs/logger"
	"route256/libs/mw/logging"
	"route256/libs/mw/metrics"
	"route256/libs/mw/recovering"
	"route256/libs/mw/tracing"
	"route256/libs/mw/validation"
	"route256/libs/tracer"
	"route256/notifications/internal/config"
	"route256/notifications/internal/domain/service"
	"route256/notifications/internal/handler"
	notificationRepo "route256/notifications/internal/repository/notification"
	"route256/notifications/internal/repository/schema"
	"route256/notifications/internal/sender"
	"route256/notifications/internal/server"
	api "route256/notifications/pkg/notifications/v1"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
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

	sender, err := sender.New(
		config.AppConfig.Telegram.Token,
		config.AppConfig.Telegram.ChatID,
		config.AppConfig.RetryCount,
	)
	if err != nil {
		log.Fatalf("failed connect to sender: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := pgxpool.Connect(ctx, config.AppConfig.DBConString)
	if err != nil {
		log.Fatalf("connect to db: %v", err)
	}

	db := dbmanager.New(pool)

	cache := cache.New[schema.Notification](
		"notifications",
		redis.NewClient(&redis.Options{
			Addr:     config.AppConfig.RedisHost,
			Password: config.AppConfig.RedisPassword, // no password set
		}),
		24*time.Hour,
	)

	repo := notificationRepo.New(db, cache)

	service := service.New(sender, repo)

	handler := handler.New(service)

	consumer, err := kafka.NewConsumer(config.AppConfig.Kafka.Brokers)
	if err != nil {
		log.Fatalf("failed connect to broker: %v", err)
	}

	receiver := broker.NewReceiver(consumer)
	receiver.RegisterHandler(config.AppConfig.Kafka.Topic, handler.HandleOrderStatus)

	notificationsServer := server.New(service)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovering.Interceptor,
			logging.Interceptor,
			validation.Interceptor,
			tracing.Interceptor,
			metrics.Interceptor,
		),
	)
	api.RegisterNotificationsServer(grpcServer, notificationsServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.AppConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go runRest(config.AppConfig, lis)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("error starting grpc")
		}
		log.Infof("Serving gRPC on %s\n", lis.Addr().String())
	}()

	if err := receiver.Subscribe(ctx, config.AppConfig.Kafka.Topic); err != nil {
		log.Fatal("failed to subscribe to topic: ", err)
	}
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

	err = api.RegisterNotificationsHandler(context.Background(), mux, conn)
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
