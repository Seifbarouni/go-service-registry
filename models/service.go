package models


type Service struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	Port      string    `json:"port"`
	Status    string    `json:"status"`
}
