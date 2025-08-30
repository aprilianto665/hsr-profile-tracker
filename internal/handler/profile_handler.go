package handler

import (
	"encoding/json"
	"fmt"
	"hsr-profile-tracker/internal/model"
	"hsr-profile-tracker/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "UID is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en", uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(model.CheckProfileResponse{
			Status:  "error",
			Message: "Failed to retrieve profile data",
			Exists:  false,
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
			Status:  "error",
			Message: "Profile not found",
			Exists:  false,
		})
	}

	return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
		Status:  "success",
		Message: "Profile exists",
		Exists:  true,
	})
}

func GetProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "UID is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en", uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve profile data",
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).Send(body)
	}

	var RawData model.RawData
	if err := json.Unmarshal(body, &RawData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse server response",
		})
	}

	util.NormalizeIconPath(&RawData.Player.Avatar.Icon)

	for i := range RawData.Characters {
		util.NormalizeIconPath(&RawData.Characters[i].Portrait)
	}

	chars := make([]model.CharacterSummary, 0, len(RawData.Characters))
	for _, c := range RawData.Characters {
		finalStats := util.MergeAttributes(c.Attributes, c.Additions)
		chars = append(chars, model.CharacterSummary{
			Name:       c.Name,
			Portrait:   c.Portrait,
			Rarity:     c.Rarity,
			Rank:       c.Rank,
			Level:      c.Level,
			Path:       c.Path,
			Element:    c.Element,
			LightCone:  c.LightCone,
			Relics:     c.Relics,
			RelicSets:  c.RelicSets,
			FinalStats: finalStats,
			RelicScore: nil,
		})
	}

	return ctx.Status(statusCode).JSON(model.APIProfileResponse{
		Status:  "success",
		Message: "Profile fetched successfully",
		Data: model.ProfileSummary{
			Player:     RawData.Player,
			Characters: chars,
		},
	})
}
