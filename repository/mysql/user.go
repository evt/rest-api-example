package mysql

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/evt/rest-api-example/mysqldb"

	"github.com/google/uuid"

	"github.com/evt/rest-api-example/model"
)

// UserMysqlRepo ...
type UserMysqlRepo struct {
	db *mysqldb.MySQL
}

// NewUserRepo ...
func NewUserRepo(db *mysqldb.MySQL) *UserMysqlRepo {
	return &UserMysqlRepo{db: db}
}

// GetUser retrieves user from Postgres
func (repo *UserMysqlRepo) GetUser(ctx context.Context, id uuid.UUID) (*model.DBUser, error) {
	if len(id) == 0 {
		return nil, errors.New("No user ID provided")
	}
	var user model.DBUser
	err := repo.db.First(&user, "id = ?", id.String()).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //not found
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates user in Postgres
func (repo *UserMysqlRepo) CreateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error) {
	if user == nil {
		return nil, errors.New("No user provided")
	}
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates user in Postgres
func (repo *UserMysqlRepo) UpdateUser(ctx context.Context, user *model.DBUser) (*model.DBUser, error) {
	err := repo.db.Save(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound { //not found
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes user in Postgres
func (repo *UserMysqlRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if len(id) == 0 {
		return errors.New("No user ID provided")
	}
	err := repo.db.Where("id = ?", id.String()).Delete(model.DBUser{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}
