package repo

import (
	"errors"
	"sync"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/api-gateway/internal/entity"
)

var (
	ErrRecordAlreadyExists = errors.New("record already exists")
	ErrRecordNotFound      = errors.New("record not found")
)

type MemoryRepo struct {
	mu   *sync.Mutex
	data map[string]*entity.UserAuth
}

func NewMemorRepo() *MemoryRepo {
	return &MemoryRepo{
		mu:   &sync.Mutex{},
		data: make(map[string]*entity.UserAuth),
	}
}

func (s *MemoryRepo) Insert(user *entity.UserAuth) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[user.Username]; exists {
		return ErrRecordAlreadyExists
	}

	s.data[user.Username] = user
	return nil
}

func (s *MemoryRepo) GetByUsername(username string) (*entity.UserAuth, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.data[username]
	if !exists {
		return nil, ErrRecordNotFound
	}
	return user, nil
}
