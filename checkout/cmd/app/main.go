package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"route256/checkout/internal/client/loms"
	"route256/checkout/internal/client/product"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain/service"
	lomspb "route256/checkout/internal/pb/loms/v1"
	productpb "route256/checkout/internal/pb/product/v1"
	cartRepo "route256/checkout/internal/repository/cart"
	"route256/checkout/internal/server"
	api "route256/checkout/pkg/checkout/v1"
	dbmanager "route256/libs/db/manager"
	log "route256/libs/logger"
	"route256/libs/mw/logging"
	"route256/libs/mw/metrics"
	"route256/libs/mw/recovering"
	"route256/libs/mw/tracing"
	"route256/libs/mw/validation"
	"route256/libs/tracer"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	if err := tracer.InitGlobal(config.AppConfig.Name); err != nil {
		log.Fatal("failed to init tracer: ", err)
	}

	lomsConn, err := getConn(
		context.Background(),
		config.AppConfig.LomsService.Addr,
		5*time.Second,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to loms server: %v", err)
	}
	defer lomsConn.Close()

	lomsClient := loms.New(lomspb.NewLomsClient(lomsConn))

	productConn, err := getConn(
		context.Background(),
		config.AppConfig.ProductsService.Addr,
		5*time.Second,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to products server: %v", err)
	}
	defer productConn.Close()

	productClient := product.New(
		productpb.NewProductServiceClient(productConn),
		config.AppConfig.ProductsService.Token,
		config.AppConfig.ProductsService.RateLimit,
		config.AppConfig.ProductsService.WorkersCount,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := pgxpool.Connect(ctx, config.AppConfig.DBConString)
	if err != nil {
		log.Fatalf("connect to db: %v", err)
	}

	db := dbmanager.New(pool)
	repo := cartRepo.New(db)

	checkoutService := service.New(
		lomsClient,
		productClient,
		repo,
	)

	checkoutServer := server.New(checkoutService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.AppConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go runRest(config.AppConfig, lis)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovering.Interceptor,
			logging.Interceptor,
			validation.Interceptor,
			tracing.Interceptor,
			metrics.Interceptor,
		),
	)
	api.RegisterCheckoutServer(grpcServer, checkoutServer)

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

	err = api.RegisterCheckoutHandler(context.Background(), mux, conn)
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

func getConn(
	ctx context.Context,
	target string,
	timeout time.Duration,
	opts ...grpc.DialOption,
) (*grpc.ClientConn, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		target,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("create client connection: %w", err)
	}

	return conn, nil
}
