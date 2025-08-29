package handler

import (
	"encoding/json"
	"fmt"
	"hsr-profile-tracker/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "uid is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en",uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(model.CheckProfileResponse{
			Status: "error",
			Message: "failed to fetch from Mihomo",
			Exists: false,
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
			Status: "error",
			Message: "profile not found",
			Exists: false,
		})
	}

	return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
		Status: "success",
		Message: "profile exists",
		Exists: true,
	})
}

func GetProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "uid is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en",uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"message": "failed to fetch from Mihomo",
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).Send(body)
	}

	var resp model.RawData
	if err := json.Unmarshal(body, &resp); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "failed to parse response",
		})
	}

	NormalizeIconPath := func (path *string) {
		const BaseIconURL = "https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/"
    	*path = BaseIconURL + *path
	}

	NormalizeIconPath(&resp.Player.Avatar.Icon)

	for i := range resp.Characters {
		NormalizeIconPath(&resp.Characters[i].Portrait)
	}

	return ctx.Status(statusCode).JSON(model.APIProfileResponse{
		Status:  "success",
		Message: "profile fetched successfully",
		Data:    model.RawData{
			Player: resp.Player,
			Characters: resp.Characters,
		},
	})
}