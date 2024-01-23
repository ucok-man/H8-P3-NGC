package repo

import (
	"errors"
	"strings"

	"github.com/ucok-man/h8-p3-ngc/03-messenger/internal/entity"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) GetByEmail(email string) (*entity.User, error) {
	User := entity.User{}

	err := s.db.Where("email = $1", email).First(&User).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &User, nil
}

func (s *UserService) GetByID(id int) (*entity.User, error) {
	User := entity.User{}

	err := s.db.Where("user_id = $1", id).First(&User).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &User, nil
}

func (s *UserService) Insert(User *entity.User) error {
	err := s.db.Create(User).Error
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			return ErrDuplicateRecord
		default:
			return err
		}
	}
	return nil
}
