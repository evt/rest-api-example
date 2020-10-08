package store

import (
	"context"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
)

// UserRepo is a store for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	GetUser(context.Context, uuid.UUID) (*model.DBUser, error)
	CreateUser(context.Context, *model.DBUser) (*model.DBUser, error)
	UpdateUser(context.Context, *model.DBUser) (*model.DBUser, error)
	DeleteUser(context.Context, uuid.UUID) error
}

// FileRepo is a store for files
//go:generate mockery --dir . --name FileRepo --output ./mocks
type FileRepo interface {
	GetFile(context.Context, uuid.UUID) (*model.DBFile, error)
	CreateFile(context.Context, *model.DBFile) (*model.DBFile, error)
	UpdateFile(context.Context, *model.DBFile) (*model.DBFile, error)
	DeleteFile(context.Context, uuid.UUID) error
}

// FileContentRepo is a store for file contennt
//go:generate mockery --dir . --name FileContentRepo --output ./mocks
type FileContentRepo interface {
	Upload(context.Context, *model.DBFile, []byte) error
	Download(context.Context, *model.DBFile) ([]byte, error)
}
