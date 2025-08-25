package handler

import "github.com/gofiber/fiber/v2"

func CheckProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Check Profile")
}

func GetProfile(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Profile")
}