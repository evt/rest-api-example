package controller

import (
	"context"
	"net/http"

	"github.com/evt/rest-api-example/logger"

	"github.com/evt/rest-api-example/lib/types"

	"github.com/evt/rest-api-example/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/evt/rest-api-example/model"
	"github.com/labstack/echo/v4"
)

// UserController ...
type UserController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewUsers creates a new user controller.
func NewUsers(ctx context.Context, services *service.Manager, logger *logger.Logger) *UserController {
	return &UserController{
		ctx:      ctx,
		services: services,
		logger:   logger,
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

	createdUser, err := ctr.services.User.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
		}
	}

	ctr.logger.Debug().Msgf("Created user '%s'", createdUser.ID.String())

	return ctx.JSON(http.StatusCreated, createdUser)
}

// Get returns user by ID
func (ctr *UserController) Get(ctx echo.Context) error {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user UUID"))
	}
	user, err := ctr.services.User.GetUser(ctx.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
	}
	return ctx.JSON(http.StatusOK, user)
}

// Delete deletes user by ID
func (ctr *UserController) Delete(ctx echo.Context) error {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user UUID"))
	}
	err = ctr.services.User.DeleteUser(ctx.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not delete user"))
		}
	}

	ctr.logger.Debug().Msgf("Deleted user '%s'", userID.String())

	return ctx.JSON(http.StatusOK, model.OK)
}
