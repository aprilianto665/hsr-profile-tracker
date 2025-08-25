package routes

import (
	"hsr-profile-tracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func ProfileRoutes(app *fiber.App){
	profile := app.Group("/")

	profile.Get("/checkprofile/:uid", handler.CheckProfile)
	profile.Get("/profile/:uid", handler.GetProfile)
}