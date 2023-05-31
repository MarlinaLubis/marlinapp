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

//@title TES SWAG
//@version 1.0
//@description This is a sample server.

//@contact.name API Support
//@contact.url https://github.com/MarlinaLubis
//@contact.email 1214040@std.ulbi.ac.id

//@host lubisapp.herokuapp.com
//@BasePath / 
//@schemes https http
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
