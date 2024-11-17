package shared

import (
	"context"

	"github.com/google/uuid"
)

type contextUserId string

const contextUserIdKey contextUserId = "user_id"

func ContextWithUserId(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, contextUserIdKey, id)
}

func UserIdFromContext(ctx context.Context) uuid.UUID {
	if id, ok := ctx.Value(contextUserIdKey).(uuid.UUID); ok {
		return id
	}
	return uuid.Nil
}
