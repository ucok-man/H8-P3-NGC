package repo

import (
	"context"
	"time"

	// "github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/contract"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	coll *mongo.Collection
}

func (s UserService) GetAll() ([]*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users = []*entity.User{}
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) Create(param *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.coll.InsertOne(ctx, param)
	if err != nil {
		return err
	}
	return nil
}
