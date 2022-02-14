package services

import (
	"errors"

	"github.com/Seifbarouni/go-service-registry/models"
)

type services []*models.Service

var allServices services = services{}
type S interface{
	GetServicesByName(name string) ([]models.Service, error)
	GetAllServices() ([]*models.Service)
	AddService(name string, ip string,port string) error
	ServiceDown(name string, ip string,port string) error
	ServiceUp(name string, ip string,port string) error
}

type service struct{}

func InitializeService() S {
	return &service{}
}

func (*service) GetAllServices() []*models.Service {
	return allServices
}

func (*service) GetServicesByName(name string) ([]models.Service, error) {
	var sv []models.Service
	for _, service := range allServices {
		if service.Name == name && service.Status == "up" {
			sv = append(sv, *service)
		}
	}
	return sv, nil
}

func (*service) AddService(name string, ip string,port string) error {
		for _, service := range allServices {
			if service.Name == name && service.Ip == ip && service.Port == port {
				return errors.New("service already exists")
			}
		}
		service := &models.Service{
			ID: len(allServices) + 1,
			Name: name,
			Ip: ip,
			Port: port,
			Status: "up",
		}
		allServices = append(allServices, service)
		return nil
}

func (*service) ServiceDown(name string, ip string,port string) error {
	for _, service := range allServices {
		if service.Name == name && service.Ip == ip && service.Port == port {
			service.Status = "down"
			return nil
		}
	}
	return errors.New("service does not exist")
}

func (*service) ServiceUp(name string, ip string,port string) error {
	for _, service := range allServices {
		if service.Name == name && service.Ip == ip && service.Port == port {
			service.Status = "up"
			return nil
		}
	}
	return errors.New("service does not exist")
}
