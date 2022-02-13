package controllers

import (
	"os"
	"sort"
	"time"

	"github.com/Seifbarouni/go-service-registry/models"
	"github.com/Seifbarouni/go-service-registry/services"

	"github.com/gofiber/fiber/v2"
)

var s services.S = services.InitializeService()

func GetService(c *fiber.Ctx) error {
	name := c.Params("serviceName")
	services, err := s.GetServicesByName(name)
	if err != nil {
		return err
	}
	return c.JSON(services)
}

func GetAllServices(c *fiber.Ctx) error {
	services, err := s.GetAllServices()
	if err != nil {
		return err
	}
	return c.JSON(services)
}

func AddService(c *fiber.Ctx) error {
	name := c.Params("serviceName")
	ip := c.Query("ip")
	if ip == "" || name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing params"})
	}
	err := s.AddService(name, ip)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(201)
}

func ServiceDown(c *fiber.Ctx) error {
	name := c.Params("serviceName")
	ip := c.Query("ip")
	if ip == "" || name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing params"})
	}
	err := s.ServiceDown(name, ip)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}

func ServiceUp(c *fiber.Ctx) error {
	name := c.Params("serviceName")
	ip := c.Query("ip")
	if ip == "" || name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing params"})
	}
	err := s.ServiceUp(name, ip)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}

func Index(c *fiber.Ctx) error {
	// get VERSION and ENV from env
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")
	if version == "" {
		version = "1.0.0"
	}
	if env == "" {
		env = "development"
	}
	// get all services
	services, err := s.GetAllServices()
	if err != nil {
		return c.Render("index", fiber.Map{
			"Version":  version,
			"Env":      env,
			"Title":    "Service Registry",
			"Date":     time.Now().Format("2006-01-02 15:04:05"),
			"Services": []models.Service{}},
		)
	}
	// sort services by name
	sort.Slice(services, func(i, j int) bool {
		return services[i].Name < services[j].Name
	})
	// for each service, if the status is down then add a class to the element
	for i := range services {
		if services[i].Status == "down" {
			services[i].Status = ""
		}
	}
	return c.Render("index", fiber.Map{
		"Version":  version,
		"Env":      env,
		"Title":    "Service Registry",
		"Date":     time.Now().Format("2006-01-02 15:04:05"),
		"Services": services,
	})
}
