package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/evt/rest-api-example/logger"

	"github.com/evt/rest-api-example/lib/types"
	"github.com/evt/rest-api-example/lib/validator"
	"github.com/evt/rest-api-example/model"
	"github.com/evt/rest-api-example/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	l := logger.Get()

	testUser := &model.User{
		Firstname: "Eugene",
		Lastname:  "Toropov",
	}

	tests := []struct {
		testName     string
		expectations func(ctx context.Context, svc *mocks.UserService)
		input        string
		err          error
		code         int
	}{
		{
			testName: "valid",
			expectations: func(ctx context.Context, svc *mocks.UserService) {
				svc.On("CreateUser", ctx, testUser).Return(testUser, nil)
			},
			input: `{ "firstname": "Eugene", "lastname": "Toropov" }`,
			code:  http.StatusCreated,
		},
		{
			testName:     "missing parameter",
			expectations: func(ctx context.Context, svc *mocks.UserService) {},
			input:        `{}`,
			err:          errors.New("code=422, message=Key: 'User.Firstname' Error:Field validation for 'Firstname' failed on the 'required' tag, Key: 'User.Lastname' Error:Field validation for 'Lastname' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			testName:     "bad request",
			expectations: func(ctx context.Context, svc *mocks.UserService) {},
			input:        `{some"}`,
			err:          errors.New("code=400, message=could not decode user data: code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string, internal=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
		{
			testName: "service error",
			expectations: func(ctx context.Context, svc *mocks.UserService) {
				svc.On("CreateUser", ctx, testUser).Return(nil, types.ErrBadRequest)
			},
			input: `{ "firstname": "Eugene", "lastname": "Toropov" }`,
			err:   errors.New("code=400, message=could not create user: bad request"),
			code:  http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Logf("running %v", test.testName)
		// initialize the echo context to use for the test
		e := echo.New()
		e.Validator = validator.NewValidator()
		r, err := http.NewRequest(echo.POST, "/users/", strings.NewReader(test.input))
		if err != nil {
			t.Fatal("could not create request")
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		svc := &mocks.UserService{}

		test.expectations(ctx.Request().Context(), svc)

		d := &UserController{ctx.Request().Context(), svc, l}
		err = d.Create(ctx)
		assert.Equal(t, test.err == nil, err == nil)
		if err != nil {
			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
			assert.Equal(t, test.code, types.HTTPCode(err))
		} else {
			assert.Equal(t, test.code, w.Code)
		}
		svc.AssertExpectations(t)
	}
}
