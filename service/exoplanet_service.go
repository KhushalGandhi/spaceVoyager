package service

import (
	"fmt"
	"spaceVoyagerProject/models"
	"spaceVoyagerProject/repository"
	"spaceVoyagerProject/utils"
)

type ExoplanetService struct {
	repo *repository.ExoplanetRepository
}

func NewExoplanetService(repo *repository.ExoplanetRepository) *ExoplanetService {
	return &ExoplanetService{repo: repo}
}

func (s *ExoplanetService) AddExoplanet(exoplanet *models.Exoplanet) *models.Exoplanet {
	return s.repo.Add(exoplanet)
}

func (s *ExoplanetService) ListExoplanets() []*models.Exoplanet {
	return s.repo.List()
}

func (s *ExoplanetService) GetExoplanetByID(id uint) *models.Exoplanet {
	return s.repo.GetByID(id)
}

func (s *ExoplanetService) UpdateExoplanet(exoplanet *models.Exoplanet) *models.Exoplanet {
	return s.repo.Update(exoplanet)
}

func (s *ExoplanetService) DeleteExoplanet(id uint) bool {
	return s.repo.Delete(id)
}

func (s *ExoplanetService) EstimateFuel(id uint, crewCapacity int) (float64, error) {
	exoplanet := s.repo.GetByID(id)
	if exoplanet == nil {
		return 0, fmt.Errorf("Exoplanet not found")
	}
	gravity := utils.CalculateGravity(exoplanet)
	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	return fuel, nil
}
