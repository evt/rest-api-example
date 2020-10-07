package pg

import (
	"context"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/db"
	"github.com/evt/rest-api-example/model"
	"github.com/go-pg/pg/v10"
)

// UserPgRepo ...
type UserPgRepo struct {
	db *db.PgDB
}

// NewUserRepo ...
func NewUserRepo(db *db.PgDB) *UserPgRepo {
	return &UserPgRepo{db: db}
}

// GetUser retrieves user from Postgres
func (repo *UserPgRepo) GetUser(ctx context.Context, id uuid.UUID) (*model.DBUser, error) {
	user := &model.DBUser{}
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

// CreateUser creates user in Postgres
func (repo *UserPgRepo) CreateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error) {
	_, err := repo.db.Model(user).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates user in Postgres
func (repo *UserPgRepo) UpdateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error) {
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

// DeleteUser deletes user in Postgres
func (repo *UserPgRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := repo.db.Model((*model.DBUser)(nil)).
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
