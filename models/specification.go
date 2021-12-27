//models/Specification.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	// Specification
	Specification struct {
		ID               uint      `gorm:"primary_key" json:"id"`
		Launched         string    `json:"launched"`
		Body             string    `json:"body"`
		OS               string    `json:"os"`
		Storage          string    `json:"storage"`
		TouchScreen      string    `json:"touch_screen"`
		SizeScreen       string    `json:"size_screen"`
		ResolutionScreen string    `json:"resolution_screen"`
		Camera           string    `json:"camera"`
		Video            string    `json:"video"`
		RAM              string    `json:"ram"`
		Chipset          string    `json:"chipset"`
		CapacityBattery  string    `json:"capacity_battery"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Devices          []Devices `json:"-"`
	}
)
