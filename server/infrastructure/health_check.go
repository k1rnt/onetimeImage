package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/k1rnt/onetimeImage/domain/model"
	"github.com/k1rnt/onetimeImage/domain/repository"
)

type healthRepository struct {
	Conn *gorm.DB
}

func NewHealthRepository(conn *gorm.DB) repository.HealthRepository {
	return &healthRepository{Conn: conn}
}

func (hr *healthRepository) Create(health *model.Health) (*model.Health, error) {
	if err := hr.Conn.Create(&health).Error; err != nil {
		return nil, err
	}

	return health, nil
}
