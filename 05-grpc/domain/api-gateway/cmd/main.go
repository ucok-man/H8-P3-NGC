package main

import (
	"log"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/app"
	_ "github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/cmd/docs"
)

// @title Hotel API
// @version 1.0
// @description Documentation for Hotel API
// @termsOfService http://swagger.io/terms/

// @contact.name ucok-man
// @contact.email ucokkocu411@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host url:8080
// @BasePath /v1
func main() {
	app := app.New()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
