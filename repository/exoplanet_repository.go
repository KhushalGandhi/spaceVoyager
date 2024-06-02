package repository

import (
	"spaceVoyagerProject/models"

	"gorm.io/gorm"
)

type ExoplanetRepository struct {
	db *gorm.DB
}

func NewExoplanetRepository(db *gorm.DB) *ExoplanetRepository {
	return &ExoplanetRepository{db: db}
}

func (r *ExoplanetRepository) Add(exoplanet *models.Exoplanet) *models.Exoplanet {
	r.db.Create(exoplanet)
	return exoplanet
}

func (r *ExoplanetRepository) List() []*models.Exoplanet {
	var exoplanets []*models.Exoplanet
	r.db.Find(&exoplanets)
	return exoplanets
}

func (r *ExoplanetRepository) GetByID(id uint) *models.Exoplanet {
	var exoplanet models.Exoplanet
	if r.db.First(&exoplanet, id).Error != nil {
		return nil
	}
	return &exoplanet
}

func (r *ExoplanetRepository) Update(exoplanet *models.Exoplanet) *models.Exoplanet {
	r.db.Save(exoplanet)
	return exoplanet
}

func (r *ExoplanetRepository) Delete(id uint) bool {
	if r.db.Delete(&models.Exoplanet{}, id).RowsAffected == 0 {
		return false
	}
	return true
}
