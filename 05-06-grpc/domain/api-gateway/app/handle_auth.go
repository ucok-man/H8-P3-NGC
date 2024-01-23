package app

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/contract"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/entity"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/jwt"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/repo"
)

func (app *Application) userLoginHandler(ctx echo.Context) error {
	var input contract.ReqUserLogin

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	user, err := app.repo.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			return app.ErrInvalidCredentials(ctx)
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	if err := user.MatchesPassword(input.Password); err != nil {
		return app.ErrInvalidCredentials(ctx)
	}

	expiration := time.Now().Add(24 * time.Hour)
	claims := jwt.NewJWTClaim(user.Username, expiration)
	token, err := jwt.GenerateToken(&claims, app.config.Jwt.Secret)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response = &contract.ResUserLogin{}
	response.AuthenticationToken.Token = token
	response.AuthenticationToken.Expiry = expiration.String()

	return ctx.JSON(http.StatusOK, response)
}

func (app *Application) userRegisterHandler(ctx echo.Context) error {
	var input contract.ReqUserRegister

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	var user entity.UserAuth
	if err := copier.Copy(&user, &input); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	if err := user.SetPassword(input.Password); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	err := app.repo.Insert(&user)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordAlreadyExists):
			return app.ErrFailedValidation(ctx, fmt.Errorf("username: already exists"))
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	var response = &contract.ResUserRegister{
		Message: "Success register",
	}
	return ctx.JSON(http.StatusAccepted, response)
}
