package controllers

import (
	s "github.com/Seifbarouni/go-service-registry/services"

	"github.com/gofiber/fiber/v2"
)

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
	err := s.AddService(name, ip)
	if err != nil {
		return c.Status(401).JSON(err)
	}
	return c.SendStatus(201)
}
