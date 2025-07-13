package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vhiweb_test/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const idleTimeout = 5 * time.Second

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	LoadEnv()

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	routers.Handle(app)

	go func() {
		if err := app.Listen(os.Getenv("APP_HOST")); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	fmt.Println("Fiber was successful shutdown.")
}
