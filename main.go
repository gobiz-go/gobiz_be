package main

import (
	// "log"

	"fmt"
	"gocroot/config"
	"log"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	
	"gocroot/url"
	"github.com/gofiber/fiber/v2"
)

func main(){
	connectdb := config.Mongoconn

	if config.ErrorMongoconn != nil {
		fmt.Println("Failed to connect to MongoDB:", config.ErrorMongoconn)
		return
	}

	// Check if the connection is successful
	if connectdb != nil {
		fmt.Println("Successfully connected to MongoDB!")
	} else {
		fmt.Println("MongoDB connection is nil")
	}

	site := fiber.New(config.Iteung)
	// site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(config.IPPort))
}
