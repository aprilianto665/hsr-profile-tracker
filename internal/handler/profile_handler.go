package handler

import (
	"encoding/json"
	"fmt"
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/model"
	"hsr-profile-tracker/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CheckProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "UID is required",
		})
	}

	if database.Rdb != nil {
		if cachedBytes, err := database.Rdb.Get(database.Ctx, uid).Bytes(); err == nil && len(cachedBytes) > 0 {
			return ctx.Status(fiber.StatusOK).JSON(model.CheckProfileResponse{
				Status:  "success",
				Message: "Profile exists",
				Exists:  true,
			})
		} else if err != nil && err != redis.Nil {
		}
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
	refresh := ctx.Query("refresh") == "true"

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "UID is required",
		})
	}

	if !refresh && database.Rdb != nil {
		if cachedBytes, err := database.Rdb.Get(database.Ctx, uid).Bytes(); err == nil && len(cachedBytes) > 0 {
			var cachedSummary model.ProfileSummary
			if err := json.Unmarshal(cachedBytes, &cachedSummary); err == nil {
				return ctx.Status(fiber.StatusOK).JSON(model.APIProfileResponse{
					Status:  "success",
					Message: "Profile fetched from cache",
					Data:    cachedSummary,
				})
			}
		} else if err != nil && err != redis.Nil {
		}
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

	player := util.NormalizePlayerAvatar(RawData.Player)

	chars := make([]model.CharacterSummary, 0, len(RawData.Characters))
	for _, c := range RawData.Characters {

		c.Path.Icon = util.NormalizeIconPath(c.Path.Icon)
		c.Element.Icon = util.NormalizeIconPath(c.Element.Icon)

		lc := util.BuildLightConeSummaryOut(c.LightCone)

		relics := make([]model.RelicSummary, 0, len(c.Relics))

		for _, r := range c.Relics {
			relics = append(relics, util.BuildRelicSummaryOut(r))
		}

		relicSets := util.NormalizeRelicSetIcons(c.RelicSets)

		finalStats := util.BuildFinalStatsOut(c.Attributes, c.Additions)

		chars = append(chars, model.CharacterSummary{
			Name:       c.Name,
			Portrait:   util.NormalizeIconPath(c.Portrait),
			Rarity:     c.Rarity,
			Rank:       c.Rank,
			Level:      c.Level,
			Path:       c.Path,
			Element:    c.Element,
			LightCone:  lc,
			Relics:     relics,
			RelicSets:  relicSets,
			FinalStats: finalStats,
			RelicScore: nil,
		})
	}

	summary := model.ProfileSummary{
		Player:     player,
		Characters: chars,
	}

	if b, err := json.Marshal(summary); err == nil {
		database.Rdb.Set(database.Ctx, uid, b, time.Hour)
	}

	return ctx.Status(statusCode).JSON(model.APIProfileResponse{
		Status:  "success",
		Message: "Profile fetched successfully",
		Data:    summary,
	})
}
