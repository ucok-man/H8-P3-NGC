package app

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/contract"
)

// users godoc
// @Tags users
// @Summary Get all users
// @Description Get all available users record
// @Accept  json
// @Produce json
// @Param sort query string false "sorting param id email nama -id -email -nama"
// @Param page query string false "current page"
// @Param page_size query string false "page size param"
// @Success 200 {object} contract.ResUserGetAll
// @Failure 422 {object} object{error=object{message=string}}
// @Failure 500 {object} object{error=object{message=string}}
// @Router /v1/users [get]
func (app *Application) userGetAllHandler(ctx echo.Context) error {
	result, err := app.gateway.User.GetAllUser()
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response contract.ResUserGetAll
	if err := copier.Copy(&response.Users, &result.Users); err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}

// users godoc
// @Tags users
// @Summary Create users
// @Description Create new users record
// @Accept  json
// @Produce json
// @Param payload body contract.ReqUserCreate true "create user"
// @Success 201 {object} contract.ResUserCreate
// @Failure 400 {object} object{error=object{message=string}}
// @Failure 422 {object} object{error=object{message=string}}
// @Failure 500 {object} object{error=object{message=string}}
// @Router /users [post]
func (app *Application) userCreateHandler(ctx echo.Context) error {
	var input contract.ReqUserCreate

	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	result, err := app.gateway.User.CreateUser(&input)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	var response contract.ResUserCreate
	if err := copier.Copy(&response.User, result.User); err != nil {
		return app.ErrInternalServer(ctx, err)
	}
	response.Status = "Success"

	return ctx.JSON(http.StatusCreated, response)
}
