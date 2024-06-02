package models

import "gorm.io/gorm"

type Exoplanet struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type"`
}
