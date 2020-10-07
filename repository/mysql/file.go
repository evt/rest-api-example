package mysql

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/evt/rest-api-example/mysqldb"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
)

// FileMysqlRepo ...
type FileMysqlRepo struct {
	db *mysqldb.MySQL
}

// NewFileRepo ...
func NewFileRepo(db *mysqldb.MySQL) *FileMysqlRepo {
	return &FileMysqlRepo{db: db}
}

// GetFile retrieves file from Postgres
func (repo *FileMysqlRepo) GetFile(ctx context.Context, id uuid.UUID) (*model.DBFile, error) {
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

// CreateFile creates file in Postgres
func (repo *FileMysqlRepo) CreateFile(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	if file == nil {
		return nil, errors.New("No file provided")
	}
	err := repo.db.Create(file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}

// UpdateFile updates file in Postgres
func (repo *FileMysqlRepo) UpdateFile(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	err := repo.db.Save(file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //not found
			return nil, nil
		}
		return nil, err
	}

	return file, nil
}

// DeleteFile deletes file in Postgres
func (repo *FileMysqlRepo) DeleteFile(ctx context.Context, id uuid.UUID) error {
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
