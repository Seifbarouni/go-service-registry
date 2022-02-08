package controllers

import (
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

func AddService(c *fiber.Ctx) error {
	name := c.Params("serviceName")
	ip := c.Query("ip")
	if ip=="" || name=="" {
		return c.Status(400).JSON(fiber.Map{"error": "missing params"})
	}
	err := s.AddService(name, ip)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(201)
}
