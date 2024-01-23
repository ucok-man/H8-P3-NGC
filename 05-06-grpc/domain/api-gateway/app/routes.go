package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/logging"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/serializer"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/validator"
)

func (app *Application) routes() http.Handler {
	router := echo.New()
	router.HTTPErrorHandler = app.httpErrorHandler
	router.JSONSerializer = serializer.JSONSerializer{}
	router.Validator = validator.New()

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.Use(app.withRecover())
	router.Use(app.withLogger())

	router.OnAddRouteHandler = func(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
		app.logger.Info("registering route", logging.Meta{
			"method": route.Method,
			"path":   route.Path,
		})
	}

	router.POST("/api/login", app.userLoginHandler)
	router.POST("/api/register", app.userRegisterHandler)

	users := router.Group("/api/users")
	users.Use(app.withLogin)
	{
		users.POST("", app.userCreateHandler)
		users.GET("", app.userGetAllHandler)
	}

	return router
}
