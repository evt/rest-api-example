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

// UserWebService ...
type UserWebService struct {
	ctx      context.Context
	userRepo repository.UserRepo
}

// NewUserWebService creates a new user web service
func NewUserWebService(
	ctx context.Context,
	user repository.UserRepo,
) *UserWebService {
	return &UserWebService{
		ctx:      ctx,
		userRepo: user,
	}
}

// GetUser ...
func (svc *UserWebService) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	userDB, err := svc.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser")
	}
	if userDB == nil {
		return nil, errors.Wrap(types.ErrBadRequest, fmt.Sprintf("User '%s' not found", userID.String()))
	}

	return userDB.ToWeb(), nil
}

// CreateUser ...
func (svc UserWebService) CreateUser(ctx context.Context, reqUser *model.User) (*model.User, error) {
	reqUser.ID = uuid.New()

	_, err := svc.userRepo.CreateUser(ctx, reqUser.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.CreateUser error")
	}

	// get created user by ID
	createdDBUser, err := svc.userRepo.GetUser(ctx, reqUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser error")
	}

	return createdDBUser.ToWeb(), nil
}

// UpdateUser ...
func (svc *UserWebService) UpdateUser(ctx context.Context, reqUser *model.User) (*model.User, error) {
	userDB, err := svc.userRepo.GetUser(ctx, reqUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser error")
	}
	if userDB == nil {
		return nil, errors.Wrap(types.ErrBadRequest, fmt.Sprintf("User '%s' not found", reqUser.ID.String()))
	}

	// update user
	_, err = svc.userRepo.UpdateUser(ctx, reqUser.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.UpdateUser error")
	}

	// get updated user by ID
	updatedDBUser, err := svc.userRepo.GetUser(ctx, reqUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser error")
	}

	return updatedDBUser.ToWeb(), nil
}

// DeleteUser ...
func (svc *UserWebService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	// Check if user exists
	userDB, err := svc.userRepo.GetUser(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.GetUser error")
	}
	if userDB == nil {
		return errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%s' not found", userID.String()))
	}

	err = svc.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser error")
	}

	return nil
}
