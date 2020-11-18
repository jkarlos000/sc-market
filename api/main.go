package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/api/conf"
	"github.com/jkarlos000/sc-market/api/router"
)

func main() {
	l := hclog.Default()
	l.Info("Iniciando API Rest")
	app := fiber.New()
	app.Use(cors.New())
	// Configurar rutas y clientes que se consumiran para comunicación RPC
	router.SetupRouter(app)
	p := conf.Config("APP_PORT")
	app.Listen(p)
}
