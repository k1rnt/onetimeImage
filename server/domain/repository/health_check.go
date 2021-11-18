package repository

import (
	"github.com/k1rnt/onetimeImage/domain/model"
)

type HealthRepository interface {
	Create(health *model.Health) (*model.Health, error)
}
