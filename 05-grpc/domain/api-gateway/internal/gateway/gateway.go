package gateway

import "github.com/ucok-man/h8-p3-ngc/05-grpc/pb"

type Service struct {
	User UserGatewayService
}

func New(userclient pb.UserServiceClient) *Service {
	return &Service{
		User: UserGatewayService{client: userclient},
	}
}
