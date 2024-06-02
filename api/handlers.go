package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"

	"spaceVoyagerProject/models"
	"spaceVoyagerProject/service"
)

func SetupRoutes(app *fiber.App, svc *service.ExoplanetService) {
	app.Post("/exoplanets", addExoplanetHandler(svc))
	app.Get("/exoplanets", listExoplanetsHandler(svc))
	app.Get("/exoplanets/:id", getExoplanetByIDHandler(svc))
	app.Put("/exoplanets/:id", updateExoplanetHandler(svc))
	app.Delete("/exoplanets/:id", deleteExoplanetHandler(svc))
	app.Get("/exoplanets/:id/fuel/:crewCapacity", estimateFuelHandler(svc))
}

func addExoplanetHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		exoplanet := new(models.Exoplanet)
		if err := c.BodyParser(exoplanet); err != nil {
			return err
		}
		exoplanet = svc.AddExoplanet(exoplanet)
		return c.JSON(exoplanet)
	}
}

func listExoplanetsHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		exoplanets := svc.ListExoplanets()
		return c.JSON(exoplanets)
	}
}

func getExoplanetByIDHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		exoplanet := svc.GetExoplanetByID(uint(id))
		if exoplanet == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exoplanet not found"})
		}
		return c.JSON(exoplanet)
	}
}

func updateExoplanetHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		exoplanet := new(models.Exoplanet)
		if err := c.BodyParser(exoplanet); err != nil {
			return err
		}
		exoplanet.ID = uint(id)
		exoplanet = svc.UpdateExoplanet(exoplanet)
		if exoplanet == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exoplanet not found"})
		}
		return c.JSON(exoplanet)
	}
}

func deleteExoplanetHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		deleted := svc.DeleteExoplanet(uint(id))
		if !deleted {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exoplanet not found"})
		}
		return c.SendString("Exoplanet deleted")
	}
}

func estimateFuelHandler(svc *service.ExoplanetService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		crewCapacity, err := strconv.Atoi(c.Params("crewCapacity"))
		if err != nil {
			return err
		}
		fuel, err := svc.EstimateFuel(uint(id), crewCapacity)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"fuel": fmt.Sprintf("%.2f units", fuel)})
	}
}
