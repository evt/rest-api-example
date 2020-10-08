package gcloud

import (
	"context"
	"fmt"

	"github.com/evt/rest-api-example/model"

	"github.com/evt/rest-api-example/lib/types"
	"github.com/evt/rest-api-example/store"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// FileContentService ...
type FileContentService struct {
	ctx   context.Context
	store *store.Store
}

// NewFileContentService creates a new file content service
func NewFileContentService(ctx context.Context, store *store.Store) *FileContentService {
	return &FileContentService{
		ctx:   ctx,
		store: store,
	}
}

// Upload file content to the cloud
func (svc *FileContentService) Upload(ctx context.Context, fileID uuid.UUID, fileBody []byte) error {
	if len(fileID) == 0 {
		return errors.New("No file provided")
	}
	// Get file from DB
	fileDB, err := svc.store.File.GetFile(ctx, fileID)
	if err != nil {
		return errors.Wrap(err, "svc.store.File.GetFile")
	}
	if fileDB == nil {
		return errors.Wrap(types.ErrBadRequest, fmt.Sprintf("File '%s' not found", fileID.String()))
	}
	// Upload file contents to the cloud
	err = svc.store.FileContent.Upload(ctx, fileDB, fileBody)
	if err != nil {
		return errors.Wrap(err, "svc.store.FileContent.Upload")
	}

	return nil
}

// Download file content from the cloud
func (svc *FileContentService) Download(ctx context.Context, fileID uuid.UUID) ([]byte, *model.DBFile, error) {
	if len(fileID) == 0 {
		return nil, nil, errors.New("No file provided")
	}
	// Get file from DB
	fileDB, err := svc.store.File.GetFile(ctx, fileID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "svc.store.File.GetFile")
	}
	if fileDB == nil {
		return nil, nil, errors.Wrap(types.ErrBadRequest, fmt.Sprintf("File '%s' not found", fileID.String()))
	}
	// Upload file contents to the cloud
	fileContent, err := svc.store.FileContent.Download(ctx, fileDB)
	if err != nil {
		return nil, nil, errors.Wrap(err, "svc.store.FileContent.Download")
	}

	return fileContent, fileDB, nil
}
