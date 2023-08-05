package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"time"
)

var (
	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "route256",
		Subsystem: "grpc",
		Name:      "histogram_response_time_seconds",
		Buckets:   []float64{},
	},
		[]string{
			"status",
			"method",
		},
	)

	counterByHandler = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "route256",
		Subsystem: "grpc",
		Name:      "counter_by_handler",
	}, []string{
		"group",
	})
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	counterByHandler.WithLabelValues(info.FullMethod).Inc()

	start := time.Now()

	h, err := handler(ctx, req)

	status := "ok"
	if err != nil {
		status = "error"
	}

	HistogramResponseTime.WithLabelValues(status, info.FullMethod).Observe(time.Since(start).Seconds())

	return h, err
}
