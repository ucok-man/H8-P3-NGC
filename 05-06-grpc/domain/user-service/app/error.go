package app

import (
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (app *Application) ErrInternal(err error, handler, context string) error {
	app.logger.Error(err, "server error", logging.Meta{
		"handler": handler,
		"context": context,
	})
	message := "the server encountered a problem and could not process your request"
	return status.Error(codes.Internal, message)
}

func (app *Application) ErrInvalidArgument(err error) error {
	return status.Error(codes.Internal, err.Error())
}

func (app *Application) ErrNotFound(ctx echo.Context, customeMsg ...string) error {
	if len(customeMsg) == 0 {
		customeMsg = append(customeMsg, "the requested resource could not be found")
	}
	message := customeMsg[0]
	return status.Error(codes.Internal, message)
}
