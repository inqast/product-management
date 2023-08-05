package validation

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator interface {
	ValidateAll() error
}

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if v, ok := req.(validator); ok {
		if validationErr := v.ValidateAll(); validationErr != nil {
			return nil, status.Errorf(codes.InvalidArgument, validationErr.Error())
		}
	}

	resp, err = handler(ctx, req)
	return resp, err
}
