package usecase

import (
	"github.com/k1rnt/onetimeImage/domain/model"
	"github.com/k1rnt/onetimeImage/domain/repository"
)

type HealthUsecase interface {
	Create(status, message string) (*model.Health, error)
}

type healthUsecase struct {
	healthRepo repository.HealthRepository
}

func NewHealthUsecase(healthRepo repository.HealthRepository) HealthUsecase {
	return &healthUsecase{healthRepo: healthRepo}
}

func (hu *healthUsecase) Create(status, message string) (*model.Health, error) {
	health, err := model.NewHealth(status, message)
	if err != nil {
		return nil, err
	}

	createdHealth, err := hu.healthRepo.Create(health)
	if err != nil {
		return nil, err
	}

	return createdHealth, nil
}
