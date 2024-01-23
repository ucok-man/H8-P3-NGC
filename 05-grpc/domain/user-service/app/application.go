package app

import (
	"sync"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/config"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/logging"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/repo"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

// const version = "1.0.0"

type Application struct {
	pb.UnimplementedUserServiceServer
	mongoclient *mongo.Client
	config      *config.Config
	logger      *logging.Logger
	repo        *repo.Services
	wg          *sync.WaitGroup
}

func New() *Application {
	logger := logging.New()
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "failed config initialization", nil)
	}

	mongoclient, err := config.OpenDB(cfg)
	if err != nil {
		logger.Fatal(err, "failed open db connection", nil)
	}
	logger.Info("database connection pool established", nil)

	app := &Application{
		wg:          new(sync.WaitGroup),
		mongoclient: mongoclient,
		logger:      logger,
		config:      cfg,
		repo:        repo.New(mongoclient.Database(cfg.Db.DBName)),
	}

	return app
}
