package users

import (
	"context"

	v1 "plants/internal/pkg/pb/users/v1"

	"plants/internal/models"
)

type storage interface {
	GetUserById(context.Context, string) (models.User, error)
	GetUserByLoginAndPassword(context.Context, string, string) (string, models.Role, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	CreateUser(context.Context, models.User) error
}

type Handler struct {
	v1.UnimplementedUserServer
	storage storage
}

func New(storage storage) *Handler {
	return &Handler{
		storage: storage,
	}
}