package util

import (
	"fmt"
	"hsr-profile-tracker/internal/model"
	"math"
)

func FormatAttributeValue(attr model.Attribute) string {
	if attr.Percent {
		return fmt.Sprintf("%.1f%%", attr.Value*100)
	}
	return fmt.Sprintf("%d", int(attr.Value))
}

func FloorToDecimal(value float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	return math.Floor(value*factor) / factor
}
