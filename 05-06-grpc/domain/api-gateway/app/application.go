package app

import (
	"sync"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/config"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/gateway"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/logging"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/repo"
)

type Application struct {
	config  *config.Config
	logger  *logging.Logger
	gateway *gateway.Service
	wg      *sync.WaitGroup
	repo    *repo.MemoryRepo
	ctxkey  struct {
		user string
	}
}

func New() *Application {
	logger := logging.New()
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "failed config initialization", nil)
	}

	userclient := cfg.InitClient()

	app := &Application{
		wg:      new(sync.WaitGroup),
		logger:  logger,
		config:  cfg,
		gateway: gateway.New(userclient),
		repo:    repo.NewMemorRepo(),
		ctxkey: struct {
			user string
		}{
			user: "user-login",
		},
	}

	return app
}
