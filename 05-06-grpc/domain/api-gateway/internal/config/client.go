package config

import (
	"fmt"
	"log"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (cfg *Config) InitClient() pb.UserServiceClient {
	connection, err := grpc.Dial(fmt.Sprintf(":%v", cfg.UserServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	return pb.NewUserServiceClient(connection)
}
