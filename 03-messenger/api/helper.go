package api

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/entity"
)

func (app *Application) getCurrentUser(ctx echo.Context) *entity.User {
	obj := ctx.Get(app.ctxkey.user)
	user, ok := obj.(*entity.User)
	if !ok {
		panic("[app.getCurrentUser]: user should be *entity.User")
	}
	return user
}

func (app *Application) getParamId(ctx echo.Context) (int, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		return int(0), fmt.Errorf("invalid id parameter")
	}

	return int(id), nil
}

func (app *Application) background(fn func()) {
	app.wg.Add(1)

	go func() {

		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.Error(fmt.Errorf("%s", err), "failed processing background task", nil)
			}
		}()

		fn()
	}()
}
