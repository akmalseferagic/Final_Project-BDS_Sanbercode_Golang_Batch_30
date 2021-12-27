//models/Comments.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	// Comments
	Comments struct {
		ID              uint      `gorm:"primary_key" json:"id"`
		Name            string    `json:"your_name"`
		CommentsSection string    `json:"comments_section"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		DevicesID       uint      `json:"devicesID"`
		Devices         Devices   `json:"-"`
		ReviewsID       uint      `json:"reviewsID"`
		Reviews         Reviews   `json:"-"`
	}
)
