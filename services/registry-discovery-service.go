package services

import (
	"errors"
	"time"

	"github.com/Seifbarouni/go-service-registry/database"
	"github.com/Seifbarouni/go-service-registry/models"
)

type S interface{
	GetServicesByName(name string) ([]models.Service, error)
	GetAllServices() ([]models.Service, error)
	AddService(name string, ip string) error
	ServiceDown(name string, ip string) error
	ServiceUp(name string, ip string) error
}

type service struct{}

func InitializeService() S {
	return &service{}
}

func (*service) GetAllServices() ([]models.Service, error) {
	var services []models.Service
	err := database.DB.Find(&services).Error
	return services, err
}

func (*service) GetServicesByName(name string) ([]models.Service, error) {
	var services []models.Service
	var upServices []models.Service= make([]models.Service, 0)
	err := database.DB.Where("name = ?", name).Find(&services).Error
	// check if service is up
	for i := range services {
		if services[i].Status == "up" {
			upServices = append(upServices, services[i])
		}
	}
	return upServices, err
}

func (*service) AddService(name string, ip string) error {
	var service models.Service
	if err := database.DB.Where("name = ? AND ip = ?", name, ip).First(&service); err.Error != nil || err.RowsAffected == 0 {
		service := models.Service{Name: name, Ip: ip, Status: "up", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		err := database.DB.Create(&service).Error
		return err
	}
	return errors.New("service already exists")
}

func (*service) ServiceDown(name string, ip string) error {
	var service models.Service
	if err := database.DB.Where("name = ? AND ip = ?", name, ip).First(&service); err.Error != nil || err.RowsAffected == 0 {
		return errors.New("service does not exist")
	}
	service.Status = "down"
	err := database.DB.Save(&service).Error
	return err
}

func (*service) ServiceUp(name string, ip string) error {
	var service models.Service
	if err := database.DB.Where("name = ? AND ip = ?", name, ip).First(&service); err.Error != nil || err.RowsAffected == 0 {
		return errors.New("service does not exist")
	}
	service.Status = "up"
	err := database.DB.Save(&service).Error
	return err
}
