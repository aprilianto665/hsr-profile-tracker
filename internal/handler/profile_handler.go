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

		formattedLightConeStats := make([]model.AttributeSummary, 0, len(c.LightCone.Attributes))

		for _, lc := range c.LightCone.Attributes {
			formattedLightConeStats = append(formattedLightConeStats, model.AttributeSummary{
				Name:  lc.Name,
				Icon:  lc.Icon,
				Value: util.FormatAttributeValue(lc),
			})
		}

		formattedLightCone := model.LightConeSummary{
			Name:       c.LightCone.Name,
			Rarity:     c.LightCone.Rarity,
			Rank:       c.LightCone.Rank,
			Level:      c.LightCone.Level,
			Icon:       c.LightCone.Icon,
			Attributes: formattedLightConeStats,
		}

		formattedRelicStats := make([]model.RelicSummary, 0, len(c.Relics))
		for _, r := range c.Relics {
			util.NormalizeIconPath(&r.Icon)
			formattedRelicStats = append(formattedRelicStats, util.BuildRelicSummaryOut(r))
		}

		finalStats := util.MergeAttributes(c.Attributes, c.Additions)
		formattedStats := make([]model.AttributeSummary, 0, len(finalStats))

		for _, fs := range finalStats {
			formattedStats = append(formattedStats, model.AttributeSummary{
				Name:  fs.Name,
				Icon:  fs.Icon,
				Value: util.FormatAttributeValue(fs),
			})
		}

		chars = append(chars, model.CharacterSummary{
			Name:       c.Name,
			Portrait:   c.Portrait,
			Rarity:     c.Rarity,
			Rank:       c.Rank,
			Level:      c.Level,
			Path:       c.Path,
			Element:    c.Element,
			LightCone:  &formattedLightCone,
			Relics:     formattedRelicStats,
			RelicSets:  c.RelicSets,
			FinalStats: formattedStats,
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
