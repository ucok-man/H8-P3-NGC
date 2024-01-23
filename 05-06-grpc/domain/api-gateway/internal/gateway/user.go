package gateway

import (
	"context"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/contract"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/pb"
)

type UserGatewayService struct {
	client pb.UserServiceClient
}

func (s UserGatewayService) CreateUser(user *contract.ReqUserCreate) (*pb.CreateUserResponse, error) {
	req := &pb.CreateUserRequest{
		User: &pb.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}
	res, err := s.client.CreateUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s UserGatewayService) GetAllUser() (*pb.GetAllUserResponse, error) {
	req := &pb.GetAllUserRequest{}
	res, err := s.client.GetAllUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
