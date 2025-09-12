package service

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func BuildProfileURL(uid string) string {
	return fmt.Sprintf("https://api.mihomo.me/sr_info_parsed/%s?lang=en", uid)
}

func HttpFetchProfile(url string) (int, []byte, error) {
	agent := fiber.Get(url).UserAgent("hsr-profile-tracker/1.0").Timeout(10 * time.Second)

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return fiber.StatusBadGateway, nil, errs[0]
	}

	return statusCode, body, nil
}
