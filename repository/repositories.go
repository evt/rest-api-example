package repository

import (
	"context"
	"github.com/google/uuid"

	"github.com/evt/simple-web-server/model"
)

// UserRepo is a repository for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	GetUser(context.Context, uuid.UUID) (*model.DBUser, error)
	CreateUser(context.Context, *model.DBUser) (*model.DBUser, error)
	UpdateUser(context.Context, *model.DBUser) (*model.DBUser, error)
	DeleteUser(context.Context, uuid.UUID) (*model.DBUser, error)
}
