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
			"message": "UID is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en",uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(model.CheckProfileResponse{
			Status: "error",
			Message: "Failed to retrieve profile data",
			Exists: false,
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
			Status: "error",
			Message: "Profile not found",
			Exists: false,
		})
	}

	return ctx.Status(statusCode).JSON(model.CheckProfileResponse{
		Status: "success",
		Message: "Profile exists",
		Exists: true,
	})
}

func GetProfile(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	if uid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "UID is required",
		})
	}

	url := fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en",uid)

	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "error",
			"message": "Failed to retrieve profile data",
		})
	}

	if statusCode < 200 || statusCode >= 300 {
		return ctx.Status(statusCode).Send(body)
	}

	var RawData model.RawData
	if err := json.Unmarshal(body, &RawData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Failed to parse server response",
		})
	}

	NormalizeIconPath := func (path *string) {
		const BaseIconURL = "https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/"
    	*path = BaseIconURL + *path
	}

	NormalizeIconPath(&RawData.Player.Avatar.Icon)

	for i := range RawData.Characters {
		NormalizeIconPath(&RawData.Characters[i].Portrait)
	}

	mergeAttributes := func(attrs, adds []model.Attribute) []model.Attribute {
		final := make([]model.Attribute, len(attrs))
		copy(final, attrs)
		idx := make(map[string]int, len(final))
		for i, a := range final {
			idx[a.Name] = i
		}
		for _, add := range adds {
			if i, ok := idx[add.Name]; ok {
				final[i].Value += add.Value
			} else {
				idx[add.Name] = len(final)
				final = append(final, add)
			}
		}
		return final
	}


	chars := make([]model.CharacterSummary, 0, len(RawData.Characters))
	for _, c := range RawData.Characters {
		finalStats := mergeAttributes(c.Attributes, c.Additions)
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
		Data:    model.ProfileSummary{
			Player:     RawData.Player,
			Characters: chars,
		},
	})
}