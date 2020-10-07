package pg

import (
	"context"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
	"github.com/evt/rest-api-example/pgdb"
	"github.com/go-pg/pg/v10"
)

// FilePgRepo ...
type FilePgRepo struct {
	db *pgdb.PgDB
}

// NewFileRepo ...
func NewFileRepo(db *pgdb.PgDB) *FilePgRepo {
	return &FilePgRepo{db: db}
}

// Get retrieves file from Postgres
func (repo *FilePgRepo) Get(ctx context.Context, id uuid.UUID) (*model.DBFile, error) {
	user := &model.DBFile{}
	err := repo.db.Model(user).
		Where("id = ?", id).
		Select()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// Create creates file in Postgres
func (repo *FilePgRepo) Create(ctx context.Context, user *model.DBFile) (*model.DBFile, error) {
	_, err := repo.db.Model(user).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates file in Postgres
func (repo *FilePgRepo) Update(ctx context.Context, user *model.DBFile) (*model.DBFile, error) {
	_, err := repo.db.Model(user).
		WherePK().
		Returning("*").
		Update()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// Delete deletes file in Postgres
func (repo *FilePgRepo) Delete(ctx context.Context, id uuid.UUID) error {
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
