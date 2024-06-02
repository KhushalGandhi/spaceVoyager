package utils

import "spaceVoyagerProject/models"

func CalculateGravity(exoplanet *models.Exoplanet) float64 {
	if exoplanet.Type == "GasGiant" {
		return 0.5 / (exoplanet.Radius * exoplanet.Radius)
	}
	return 1 / (exoplanet.Radius * exoplanet.Radius)
}
