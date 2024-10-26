package main

import (
	"fmt"
	"go-hexagonal-example/api"
	"go-hexagonal-example/config"
	"go-hexagonal-example/repository"
	"go-hexagonal-example/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Use(cache.New())
	app.Use(recover.New())

	repo, _ := repository.NewMongoRepository(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB_NAME"), 1000)
	service := service.NewProductService(repo)
	handler := api.NewHandler(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3011"
	}

	app.Get("/products/{code}", handler.Get)
	app.Post("/products", handler.Post)
	app.Delete("/products/:id", handler.Delete)
	app.Get("/products", handler.GetAll)
	app.Put("/products", handler.Put)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
