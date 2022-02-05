package models

import "time"

type Service struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
