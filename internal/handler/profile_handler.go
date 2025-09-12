package handler

import (
	"encoding/json"
	"hsr-profile-tracker/internal/model"
	"hsr-profile-tracker/internal/service"
	"hsr-profile-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
)

func GetUID(ctx *fiber.Ctx) (string, bool) {
	uid := ctx.Params("uid")

	if uid == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "UID is required",
		})
		return "", true
	}

	return uid, false
}

func CheckProfile(ctx *fiber.Ctx) error {
	uid, bad := GetUID(ctx)
	if bad {
		return nil
	}

	if _, ok := service.CacheGetSummary(uid); ok {
		return ctx.Status(fiber.StatusOK).JSON(model.CheckProfileResponse{
			Status:  "success",
			Message: "Profile exists",
			Exists:  true,
		})
	}

	url := service.BuildProfileURL(uid)

	statusCode, _, errs := service.HttpFetchProfile(url)
	if errs != nil {
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
	uid, bad := GetUID(ctx)
	if bad {
		return nil
	}

	refresh := ctx.Query("refresh") == "true"

	if !refresh {
		if cachedSummary, ok := service.CacheGetSummary(uid); ok {
			return ctx.Status(fiber.StatusOK).JSON(model.APIProfileResponse{
				Status:  "success",
				Message: "Profile fetched from cache",
				Data:    *cachedSummary,
			})
		}
	}

	url := service.BuildProfileURL(uid)

	statusCode, body, errs := service.HttpFetchProfile(url)
	if errs != nil {
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

	summary := util.BuildProfileSummaryOut(RawData)

	if b, err := json.Marshal(summary); err == nil {
		service.CacheSetBytes(uid, b)
	}

	return ctx.Status(statusCode).JSON(model.APIProfileResponse{
		Status:  "success",
		Message: "Profile fetched successfully",
		Data:    summary,
	})
}
