package handler

import "github.com/gofiber/fiber/v2"

func CheckProfile(ctx *fiber.Ctx){
	ctx.SendString("Check Profile")
}