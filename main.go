package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"spaceVoyagerProject/api"
	"spaceVoyagerProject/config"
	"spaceVoyagerProject/db"
	"spaceVoyagerProject/models"
	"spaceVoyagerProject/repository"
	"spaceVoyagerProject/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize GORM with PostgreSQL
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Migrate the schema
	if err := database.AutoMigrate(&models.Exoplanet{}); err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	app := fiber.New()

	// Initialize repository
	repo := repository.NewExoplanetRepository(database)

	// Initialize service
	svc := service.NewExoplanetService(repo)

	// Initialize handlers
	api.SetupRoutes(app, svc)

	log.Fatal(app.Listen(":3000"))
}
