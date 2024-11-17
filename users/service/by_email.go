package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/prathoss/telemetry_showcase/proto/users"
	"github.com/prathoss/telemetry_showcase/shared"
	"github.com/prathoss/telemetry_showcase/users/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetUserByEmail(ctx context.Context, request *users.GetUserByEmailRequest) (*users.UserReply, error) {
	val, err := s.cache.Get(
		ctx,
		5*time.Minute,
		fmt.Sprintf("user:%s", request.Email),
		func(ctx context.Context, key string) (string, error) {
			u, err := s.repository.GetUserByEmail(ctx, request.Email)
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
		slog.WarnContext(ctx, "user not found", slog.String("email", request.Email))
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
