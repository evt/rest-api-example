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
	DeleteUser(context.Context, uuid.UUID) error
}

// FileRepo is a repository for files
//go:generate mockery --dir . --name FileRepo --output ./mocks
type FileRepo interface {
	Get(context.Context, uuid.UUID) (*model.DBFile, error)
	Create(context.Context, *model.DBFile) (*model.DBFile, error)
	Update(context.Context, *model.DBFile) (*model.DBFile, error)
	Delete(context.Context, uuid.UUID) error
}

// FileContentRepo is a repository for file contennt
//go:generate mockery --dir . --name FileContentRepo --output ./mocks
type FileContentRepo interface {
	Upload(context.Context, *model.DBFile, []byte) error
	Download(context.Context, *model.DBFile) ([]byte, error)
}
