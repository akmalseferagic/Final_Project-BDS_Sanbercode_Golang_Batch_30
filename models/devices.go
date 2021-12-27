//models/Devices.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	// Devices
	Devices struct {
		ID              uint          `gorm:"primary_key" json:"id"`
		Name            string        `json:"name"`
		Description     string        `json:"description"`
		Price           string        `json:"price"`
		CreatedAt       time.Time     `json:"created_at"`
		UpdatedAt       time.Time     `json:"updated_at"`
		Reviews         []Reviews     `json:"-"`
		BrandID         uint          `json:"brandID"`
		Brand           Brand         `json:"-"`
		SpecificationID uint          `json:"specificationID"`
		Specification   Specification `json:"-"`
	}
)
