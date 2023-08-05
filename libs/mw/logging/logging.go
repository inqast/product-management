package logging

import (
	"context"
	log "route256/libs/logger"

	"google.golang.org/grpc"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		log.Errorf(ctx, info.FullMethod, "error while processing handler, err=%s", err)
	}

	return resp, err
}
