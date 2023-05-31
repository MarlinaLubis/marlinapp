package main

import (
	"log"

	"github.com/MarlinaLubis/marlinapp/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/MarlinaLubis/marlinapp/url"

	"github.com/gofiber/fiber/v2"
	
	_ "github.com/MarlinaLubis/marlinapp/docs"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
