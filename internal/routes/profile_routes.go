package routes

import "github.com/gofiber/fiber/v2"

func ProfileRoutes(app *fiber.App){
	profile := app.Group("profile")

	profile.Get("/checkprofile/:uid", func (ctx *fiber.Ctx) error {})
	profile.Get("/profile/:uid", func (ctx *fiber.Ctx) error {})
}