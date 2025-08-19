package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"type:uint;primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
}
