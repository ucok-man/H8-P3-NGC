package app

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/entity"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/pb"
)

func (app *Application) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var user entity.User
	if err := copier.Copy(&user, input.User); err != nil {
		return nil, app.ErrInternal(err, "CreateUser", "copy request to user entity")
	}

	if user.ID == "" || user.Name == "" {
		return nil, app.ErrInvalidArgument(fmt.Errorf("invalid user.id or user.name parameter"))
	}

	if err := app.repo.User.Create(&user); err != nil {
		return nil, app.ErrInternal(err, "CreateUser", "inserting record to mongodb")
	}

	var response = &pb.CreateUserResponse{
		User: &pb.User{},
	}
	if err := copier.Copy(&response.User, &user); err != nil {
		return nil, app.ErrInternal(err, "CreateUser", "copy user to response")
	}

	return response, nil
}

func (app *Application) GetAllUser(ctx context.Context, input *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	users, err := app.repo.User.GetAll()
	if err != nil {
		return nil, app.ErrInternal(err, "GetAllUser", "get all record from mongodb")
	}

	var response = &pb.GetAllUserResponse{}
	if err := copier.Copy(&response.Users, &users); err != nil {
		return nil, app.ErrInternal(err, "GetAllUser", "copy users to response")
	}
	return response, nil
}
