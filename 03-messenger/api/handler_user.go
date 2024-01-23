package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/contract"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/entity"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/jwt"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/repo"
)

// users godoc
// @Tags users
// @Summary Create user
// @Description Create new user record
// @Accept  json
// @Produce json
// @Param payload body contract.ReqUserRegister true "Create User"
// @Success 201 {object} contract.ResUserRegister
// @Failure 400 {object} object{error=object{message=string}}
// @Failure 422 {object} object{error=object{message=string}}
// @Failure 500 {object} object{error=object{message=string}}
// @Router /users/register [post]
func (app *Application) userRegisterHandler(ctx echo.Context) error {
	var input contract.ReqUserRegister

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	var user entity.User
	if err := copier.Copy(&user, &input); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	if err := user.SetPassword(input.Password); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	err := app.repo.User.Insert(&user)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicateRecord):
			return app.ErrFailedValidation(ctx, fmt.Errorf("email: already exists"))
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	var response contract.ResUserRegister
	if err := copier.Copy(&response.User, user); err != nil {
		return app.ErrInternalServer(ctx, err)
	}
	response.Message = "success"

	return ctx.JSON(http.StatusAccepted, response)
}

// users godoc
// @Tags users
// @Summary Login user
// @Description Login user record
// @Accept  json
// @Produce json
// @Param payload body contract.ReqUserLogin true "Login User"
// @Success 200 {object} contract.ResUserLogin
// @Failure 400 {object} object{error=object{message=string}}
// @Failure 401 {object} object{error=object{message=string}}
// @Failure 422 {object} object{error=object{message=string}}
// @Failure 500 {object} object{error=object{message=string}}
// @Router /users/login [post]
func (app *Application) userLoginHandler(ctx echo.Context) error {
	var input contract.ReqUserLogin

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	user, err := app.repo.User.GetByEmail(input.Email)
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
	claims := jwt.NewJWTClaim(user.UserID, expiration)
	token, err := jwt.GenerateToken(&claims, app.config.Jwt.Secret)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response = &contract.ResUserLogin{}
	response.AuthenticationToken.Token = token
	response.AuthenticationToken.Expiry = expiration.String()

	return ctx.JSON(http.StatusOK, response)
}
