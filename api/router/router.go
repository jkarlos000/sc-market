package router

import (
	"github.com/gofiber/fiber"
	"github.com/jkarlos000/sc-market/api/handlers"
	"net/http"
)

type msg struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func SetupRouter(app *fiber.App)  {
	// Middleware
	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) {
		ctx.Status(http.StatusMethodNotAllowed)
		_ = ctx.JSON(msg{
			Message: "No permitido",
			Status:  http.StatusMethodNotAllowed,
		})
	})

	// Login
	login := api.Group("/login")
	loginHandler := handlers.NewLogin()
	loginHandler.SetupRoutes(login)


}
