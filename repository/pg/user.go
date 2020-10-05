package pg

import (
	"context"
	"github.com/google/uuid"

	"github.com/evt/simple-web-server/db"
	"github.com/evt/simple-web-server/model"
	"github.com/go-pg/pg/v10"
)

// UserPgRepo
type UserPgRepo struct {
	db *db.PgDB
}

// NewUserRepo
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
func (repo *UserPgRepo) DeleteUser(ctx context.Context, id uuid.UUID) (*model.DBUser, error) {
	user := &model.DBUser{}
	_, err := repo.db.Model(user).
		Where("id = ?", id).
		Returning("*").
		Delete()
	if err != nil {
		if err == pg.ErrNoRows { //not found
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
