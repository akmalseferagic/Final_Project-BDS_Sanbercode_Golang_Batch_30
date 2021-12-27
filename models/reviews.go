//models/Reviews.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	// Reviews
	Reviews struct {
		ID        uint       `gorm:"primary_key" json:"id"`
		Name      string     `json:"name_article"`
		Article   string     `json:"article"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Comments  []Comments `json:"-"`
		Opinions  []Opinions `json:"-"`
		DevicesID uint       `json:"devicesID"`
		Devices   Devices    `json:"-"`
	}
)
