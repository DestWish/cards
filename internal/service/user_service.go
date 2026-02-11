package service

import (
	"errors"

	"github.com/DestWish/cards/internal/models"
	"github.com/DestWish/cards/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Username: req.Username,
		Password: HashPassword(req.Password),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(p string) string {
	return p
}

func (s *UserService) UpdateUserPassword(req models.UpdateUserPassword) (*models.User, error) {
	user, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, errors.New("User not found")
	}
	user.Password = HashPassword(req.Password)

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
