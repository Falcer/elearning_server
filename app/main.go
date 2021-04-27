package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"server/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	port string
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Error load .env")
	}
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		host, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": fmt.Sprintf("App running 🔥 at %s", host),
		})
	})
	_ = auth.NewRouter(app)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen("0.0.0.0:8080"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
