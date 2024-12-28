package main

import (
	"log"

	"github.com/Nidasakinaa/ws-KaloriKu/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Nidasakinaa/ws-KaloriKu/url"

	_ "github.com/Nidasakinaa/ws-KaloriKu/docs"
	"github.com/gofiber/fiber/v2"
)

// @title TES SWAGGER KALORIKU
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Nidasakinaa
// @contact.email 714220040@std.ulbi.ac.id

// @host kaloriku-d9941de09573.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
