package pg

import (
	"context"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
	"github.com/go-pg/pg/v10"
)

// FilePgRepo ...
type FilePgRepo struct {
	db *DB
}

// NewFileMetaRepo ...
func NewFileMetaRepo(db *DB) *FilePgRepo {
	return &FilePgRepo{db: db}
}

// GetFileMeta retrieves file from MySQL
func (repo *FilePgRepo) GetFileMeta(ctx context.Context, id uuid.UUID) (*model.DBFile, error) {
	file := &model.DBFile{}
	err := repo.db.Model(file).
		Where("id = ?", id).
		Select()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}
	return file, nil
}

// CreateFileMeta creates file in Postgres
func (repo *FilePgRepo) CreateFileMeta(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	_, err := repo.db.Model(file).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return file, nil
}

// UpdateFileMeta updates file in Postgres
func (repo *FilePgRepo) UpdateFileMeta(ctx context.Context, file *model.DBFile) (*model.DBFile, error) {
	_, err := repo.db.Model(file).
		WherePK().
		Returning("*").
		Update()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}

	return file, nil
}

// DeleteFileMeta deletes file in Postgres
func (repo *FilePgRepo) DeleteFileMeta(ctx context.Context, id uuid.UUID) error {
	_, err := repo.db.Model((*model.DBFile)(nil)).
		Where("id = ?", id).
		Delete()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}
