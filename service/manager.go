package service

import (
	"context"

	"github.com/evt/rest-api-example/service/gcloud"

	"github.com/evt/rest-api-example/service/web"

	"github.com/evt/rest-api-example/store"
	"github.com/pkg/errors"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	User        UserService
	File        FileMetaService
	FileContent FileContentService
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		User:        web.NewUserWebService(ctx, store),
		File:        web.NewFileWebService(ctx, store),
		FileContent: gcloud.NewFileContentService(ctx, store),
	}, nil
}
