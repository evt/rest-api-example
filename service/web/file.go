package web

import (
	"context"
	"fmt"

	"github.com/evt/simple-web-server/lib/types"
	"github.com/evt/simple-web-server/model"
	"github.com/evt/simple-web-server/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// FileWebService ...
type FileWebService struct {
	ctx      context.Context
	fileRepo repository.FileRepo
}

// NewFileWebService creates a new file web service
func NewFileWebService(ctx context.Context, file repository.FileRepo) *FileWebService {
	return &FileWebService{
		ctx:      ctx,
		fileRepo: file,
	}
}

// GetFile ...
func (svc *FileWebService) GetFile(ctx context.Context, fileID uuid.UUID) (*model.File, error) {
	fileDB, err := svc.fileRepo.Get(ctx, fileID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.GetFile")
	}
	if fileDB == nil {
		return nil, errors.Wrap(types.ErrBadRequest, fmt.Sprintf("File '%s' not found", fileID.String()))
	}

	return fileDB.ToWeb(), nil
}

// CreateFile ...
func (svc FileWebService) CreateFile(ctx context.Context, reqFile *model.File) (*model.File, error) {
	reqFile.ID = uuid.New()

	_, err := svc.fileRepo.Create(ctx, reqFile.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.CreateFile error")
	}

	// get created file by ID
	createdDBFile, err := svc.fileRepo.Get(ctx, reqFile.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.GetFile error")
	}

	return createdDBFile.ToWeb(), nil
}

// UpdateFile ...
func (svc *FileWebService) UpdateFile(ctx context.Context, reqFile *model.File) (*model.File, error) {
	fileDB, err := svc.fileRepo.Get(ctx, reqFile.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.GetFile error")
	}
	if fileDB == nil {
		return nil, errors.Wrap(types.ErrBadRequest, fmt.Sprintf("File '%s' not found", reqFile.ID.String()))
	}

	// update file
	_, err = svc.fileRepo.Update(ctx, reqFile.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.UpdateFile error")
	}

	// get updated file by ID
	updatedDBFile, err := svc.fileRepo.Get(ctx, reqFile.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.file.GetFile error")
	}

	return updatedDBFile.ToWeb(), nil
}

// DeleteFile ...
func (svc *FileWebService) DeleteFile(ctx context.Context, fileID uuid.UUID) error {
	// Check if file exists
	fileDB, err := svc.fileRepo.Get(ctx, fileID)
	if err != nil {
		return errors.Wrap(err, "svc.file.GetFile error")
	}
	if fileDB == nil {
		return errors.Wrap(types.ErrNotFound, fmt.Sprintf("File '%s' not found", fileID.String()))
	}

	err = svc.fileRepo.Delete(ctx, fileID)
	if err != nil {
		return errors.Wrap(err, "svc.file.DeleteFile error")
	}

	return nil
}
