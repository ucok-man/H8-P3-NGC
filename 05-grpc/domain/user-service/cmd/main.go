package main

import (
	"log"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/app"
)

func main() {
	if err := app.New().Serve(); err != nil {
		log.Fatal(err)
	}
}
