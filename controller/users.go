package controller

import (
	"context"
	"github.com/evt/simple-web-server/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"

	"github.com/evt/simple-web-server/model"
	"github.com/labstack/echo/v4"
)

// Users controller
type UserController struct {
	ctx     context.Context
	userSvc service.UserService
}

// NewUsers creates a new user controller.
func NewUsers(ctx context.Context, userSvc service.UserService) *UserController {
	return &UserController{
		ctx:     ctx,
		userSvc: userSvc,
	}
}

// Create creates new user
func (ctr *UserController) Create(ctx echo.Context) error {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}
	err = ctx.Validate(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	_, err = ctr.userSvc.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// Get fetches user from DB
func (ctr *UserController) Get(ctx echo.Context) error {
	userID, err := uuid.Parse(ctx.QueryParam("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user UUID"))
	}
	user, err := ctr.userSvc.GetUser(ctx.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not read user"))
	}
	return ctx.JSON(http.StatusOK, user)
}
