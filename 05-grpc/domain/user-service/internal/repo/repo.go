package repo

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrRecordNotFound  = errors.New("no record found")
	ErrDuplicateRecord = errors.New("record duplicate on unique constraint")
)

const (
	TransactionsColl = "transactions"
	UserColl         = "users"
	RoomColl         = "rooms"
)

type Services struct {
	User UserService
}

func New(db *mongo.Database) *Services {
	return &Services{
		User: UserService{coll: db.Collection(UserColl)},
	}
}
