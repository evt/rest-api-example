package mysql

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
)

// FileMysqlRepo ...
type FileMysqlRepo struct {
	db *MySQL
}

// NewFileMetaRepo ...
func NewFileMetaRepo(db *MySQL) *FileMysqlRepo {
	return &FileMysqlRepo{db: db}
}

// GetFileMeta retrieves file from Postgres
func (repo *FileMysqlRepo) GetFileMeta(ctx context.Context, id uuid.UUID) (*model.DBFile, error) {
	if len(id) == 0 {
		return nil, errors.New("No file ID provided")
	}
	var file model.DBFile
	err := repo.db.First(&file, "id = ?", id.String()).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //not found
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

// CreateFileMeta creates file in Postgres
func (repo *FileMysqlRepo) CreateFileMeta(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	if file == nil {
		return nil, errors.New("No file provided")
	}
	err := repo.db.Create(file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}

// UpdateFileMeta updates file in Postgres
func (repo *FileMysqlRepo) UpdateFileMeta(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	err := repo.db.Save(file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //not found
			return nil, nil
		}
		return nil, err
	}

	return file, nil
}

// DeleteFileMeta deletes file in Postgres
func (repo *FileMysqlRepo) DeleteFileMeta(ctx context.Context, id uuid.UUID) error {
	if len(id) == 0 {
		return errors.New("No file ID provided")
	}
	err := repo.db.Where("id = ?", id.String()).Delete(model.DBFile{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}
