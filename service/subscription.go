package service

import (
	"subscription-service/model"
	"subscription-service/repository"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepo
}

func NewSubscriptionService(r *repository.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{repo: r}
}

func (s *SubscriptionService) Create(sub *model.Subscription) error {
	return s.repo.Create(sub)
}

func (s *SubscriptionService) GetAll() ([]model.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetByID(id int) (*model.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *SubscriptionService) Update(id int, sub *model.Subscription) error {
	return s.repo.Update(id, sub)
}

func (s *SubscriptionService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *SubscriptionService) Total(service, userId string) (float64, error) {
	return s.repo.GetTotal(service, userId)
}
