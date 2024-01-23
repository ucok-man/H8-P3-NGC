package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/jwt"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/logging"
	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/repo"
)

func (app *Application) withRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			app.logger.Error(err, "PANIC RECOVER", logging.Meta{
				"stack": string(stack),
			})
			return err
		},
	})
}

func (app *Application) withLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:     true,
		LogStatus:       true,
		LogMethod:       true,
		LogURI:          true,
		LogLatency:      true,
		LogResponseSize: true,
		LogError:        true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			switch v.Status {
			case http.StatusInternalServerError:
				httperr, _ := v.Error.(*echo.HTTPError)
				app.logger.Error(httperr.Internal, http.StatusText(500), logging.Meta{
					"code":          v.Status,
					"method":        v.Method,
					"url":           v.URI,
					"ip_addr":       v.RemoteIP,
					"response_time": v.Latency,
					"response_size": v.ResponseSize,
					"stack":         v.Error,
				})
			default:
				app.logger.Info(http.StatusText(v.Status), logging.Meta{
					"code":          v.Status,
					"method":        v.Method,
					"url":           v.URI,
					"ip_addr":       v.RemoteIP,
					"response_time": v.Latency,
					"response_size": v.ResponseSize,
				})
			}

			return nil
		},
	})
}

func (app *Application) withLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := ctx.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return app.ErrInvalidAuthenticationToken(ctx)
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return app.ErrInvalidAuthenticationToken(ctx)
		}
		tokenstr := headerParts[1]

		var claim jwt.JWTClaim
		err := jwt.DecodeToken(tokenstr, &claim, app.config.Jwt.Secret)
		if err != nil {
			return app.ErrInvalidAuthenticationToken(ctx)
		}

		user, err := app.repo.User.GetByID(claim.UserID)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrRecordNotFound):
				return app.ErrInvalidAuthenticationToken(ctx)
			default:
				return app.ErrInternalServer(ctx, err)
			}
		}

		// set context
		ctx.Set(app.ctxkey.user, user)
		return next(ctx)
	}
}
