package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/users"
	"github.com/prathoss/telemetry_showcase/shared"
	"github.com/prathoss/telemetry_showcase/users/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetUserById(ctx context.Context, request *users.GetUserByIdRequest) (*users.UserReply, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	slog.InfoContext(ctx, "get user by id", slog.String("id", id.String()))

	val, err := s.cache.Get(
		ctx,
		5*time.Minute,
		fmt.Sprintf("user:%s", id.String()),
		func(ctx context.Context, key string) (string, error) {
			u, err := s.repository.GetUserByID(ctx, id)
			if err != nil {
				return "", err
			}
			val, err := json.Marshal(u)
			if err != nil {
				return "", err
			}
			return string(val), nil
		},
	)
	var errNotFound *shared.ErrNotFound
	if errors.As(err, &errNotFound) {
		slog.WarnContext(ctx, "user not found", slog.String("id", id.String()))
		return nil, status.Error(codes.NotFound, errNotFound.Error())
	}
	if err != nil {
		return nil, err
	}

	var user dao.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return &users.UserReply{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
