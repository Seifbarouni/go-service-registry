package models

import "time"

type Service struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
