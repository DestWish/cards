package service

import (
	"errors"

	"github.com/DestWish/cards/internal/models"
	"github.com/DestWish/cards/internal/repository"
)

type CardService struct {
	repo     *repository.CardRepository
	userRepo *repository.UserRepository
}

func NewCardService(repo *repository.CardRepository, userRepo *repository.UserRepository) *CardService {
	return &CardService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *CardService) CreateCard(req models.CreateCardRequest) (*models.Card, error) {
	if _, err := s.userRepo.GetById(req.UserID); err != nil {
		return nil, errors.New("User not found")
	}

	card := &models.Card{
		UserID:   req.UserID,
		Topic:    req.Topic,
		Question: req.Question,
		Answer:   req.Answer,
	}
	if err := s.repo.Create(card); err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) UpdateCard(req models.UpdateCardRequest) (*models.Card, error) {
	card, err := s.repo.GetById(req.CardID)
	if err != nil {
		return nil, errors.New("Card not found")
	}
	if req.Topic != nil {
		card.Topic = *req.Topic
	}
	if req.Question != nil {
		card.Question = *req.Question
	}
	if req.Answer != nil {
		card.Answer = *req.Answer
	}
	if err := s.repo.Update(card); err != nil {
		return nil, err
	}
	return card, nil
}

func (s *CardService) GetAllCards(userID uint) ([]models.Card, error) {
	return s.repo.GetAllByUser(userID)
}
