package repo

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound    = errors.New("no record found")
	ErrDuplicateRecord   = errors.New("record duplicate on unique constraint")
	ErrViolateForeignKey = errors.New("record violate foreign key constraint")
	ErrEditConflict      = errors.New("record edit conflict when updating")
)

type Services struct {
	User UserService
}

func New(db *gorm.DB) *Services {
	return &Services{
		User: UserService{db: db},
	}
}
