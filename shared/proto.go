package shared

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func GrpcLoggingInterceptor(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, request)
	if err != nil {
		slog.ErrorContext(ctx, "request failed", Err(err), "took", time.Since(start), "method", info.FullMethod)
	} else {
		slog.InfoContext(ctx, "request succeeded", "took", time.Since(start), "method", info.FullMethod)
	}
	return resp, err
}
