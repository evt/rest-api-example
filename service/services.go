package service

import (
	"context"

	"github.com/evt/simple-web-server/model"
	"github.com/google/uuid"
)

// UserService is a service for users
//go:generate mockery --dir . --name UserService --output ./mocks
type UserService interface {
	GetUser(context.Context, uuid.UUID) (*model.User, error)
	CreateUser(context.Context, *model.User) (*model.User, error)
	UpdateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, uuid.UUID) error
}
