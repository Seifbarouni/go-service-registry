package services

import (
	"errors"
	"time"

	"github.com/Seifbarouni/go-service-registry/database"
	"github.com/Seifbarouni/go-service-registry/models"
)

func GetServicesByName(name string) ([]models.Service, error) {
	var services []models.Service
	err := database.DB.Where("name = ?", name).Find(&services).Error
	return services, err
}

func AddService(name string, ip string) error {
	var service models.Service
	if err := database.DB.Where("name = ? AND ip = ?", name, ip).First(&service); err.Error != nil || err.RowsAffected == 0 {
		service := models.Service{Name: name, Ip: ip, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		err := database.DB.Create(&service).Error
		return err
	}
	return errors.New("service already exists")
}
