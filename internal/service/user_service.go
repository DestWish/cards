package service

import (
	"errors"

	"github.com/DestWish/cards/internal/models"
	"github.com/DestWish/cards/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	hash, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: hash,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUserPassword(req models.UpdateUserPassword) (*models.User, error) {
	user, err := s.repo.GetById(req.ID)
	if err != nil {
		return nil, errors.New("User not found")
	}
	user.Password, err = HashPassword(req.Password)

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(bytes), err
}

func IsPasswordCorrect(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
