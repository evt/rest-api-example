package service

import (
	"context"

	"github.com/evt/rest-api-example/model"
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

// FileMetaService is a service for files
//go:generate mockery --dir . --name FileMetaService --output ./mocks
type FileMetaService interface {
	GetFileMeta(context.Context, uuid.UUID) (*model.File, error)
	CreateFileMeta(context.Context, *model.File) (*model.File, error)
	UpdateFileMeta(context.Context, *model.File) (*model.File, error)
	DeleteFileMeta(context.Context, uuid.UUID) error
}

// FileContentService is a service to upload file content
//go:generate mockery --dir . --name FileContentService --output ./mocks
type FileContentService interface {
	Upload(context.Context, uuid.UUID, []byte) error
	Download(context.Context, uuid.UUID) ([]byte, *model.DBFile, error)
}
