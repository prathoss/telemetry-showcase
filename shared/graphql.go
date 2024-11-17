package shared

import (
	"context"
	"log/slog"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

var _ graphql.ResponseInterceptor = (*GraphqlLogger)(nil)
var _ graphql.HandlerExtension = (*GraphqlLogger)(nil)

type GraphqlLogger struct {
}

func (g *GraphqlLogger) ExtensionName() string {
	return "GraphqlLogger"
}

func (g *GraphqlLogger) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (g *GraphqlLogger) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()
	resp := next(ctx)
	elapsed := time.Since(start)
	slog.InfoContext(ctx, "graphql request", "elapsed", elapsed, "errors", resp.Errors.Error())
	return resp
}
