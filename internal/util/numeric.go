package util

import (
	"hsr-profile-tracker/internal/model"
	"math"
)

func FormatAttributeValue(attr model.Attribute) float64 {
	if attr.Percent {
		return FloorToDecimal(attr.Value*100, 2)
	}
	return FloorToDecimal(attr.Value, 0)
}

func FloorToDecimal(value float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	return math.Floor(value*factor) / factor
}
