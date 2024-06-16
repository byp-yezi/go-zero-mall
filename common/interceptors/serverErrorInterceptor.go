package interceptors

import (
	"context"
	"go-zero-mall/common/codex"

	"google.golang.org/grpc"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		return resp, codex.GetGrpcStatusFromError(err).Err()
	}
}
