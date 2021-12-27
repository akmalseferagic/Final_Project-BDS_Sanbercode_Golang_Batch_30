//models/Opinions.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	// Opinions
	Opinions struct {
		ID             uint      `gorm:"primary_key" json:"id"`
		Name           string    `json:"your_name"`
		OpinionSection string    `json:"opinion_section"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		ReviewsID      uint      `json:"reviewsID"`
		Reviews        Reviews   `json:"-"`
	}
)
