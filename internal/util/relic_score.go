package util

import (
	"fmt"
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/model"
	"math"
	"strconv"
)

func FindCharacterWeights(player model.Player, char model.Character) model.CharacterWeights {
	if char.Name == player.Nickname {
		trailblazer := fmt.Sprintf("Trailblazer (%s)", char.Element.Name)

		result := configs.CharacterWeights[trailblazer]

		return result
	}

	result := configs.CharacterWeights[char.Name]

	return result
}

func FindStatCoefficient(stat string) float64 {
	return configs.StatWeights.CoefficientStat[stat]
}

func CalculateMainStatScore(r model.Relic, charWeight model.CharacterWeights) float64 {
	slotMap := map[int]string{
		3: "Body",
		4: "Feet",
		5: "Sphere",
		6: "Rope",
	}

	slotName := slotMap[r.Type]
	if slotName == "" {
		return 0
	}

	recommendedStats := charWeight.MainStats
	var isRecommended bool

	switch slotName {
	case "Body":
		isRecommended = contains(recommendedStats.Body, r.MainAffix.Type)
	case "Feet":
		isRecommended = contains(recommendedStats.Feet, r.MainAffix.Type)
	case "Sphere":
		isRecommended = contains(recommendedStats.Sphere, r.MainAffix.Type)
	case "Rope":
		isRecommended = contains(recommendedStats.Rope, r.MainAffix.Type)
	}

	if isRecommended {
		return 5.832
	}

	return 0
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CalculateRelicScoreValue(r model.Relic, player model.Player, char model.Character) float64 {
	charWeight := FindCharacterWeights(player, char)
	partWeight := charWeight.RelicPartWeight[strconv.Itoa(r.Type)]

	var totalScore float64

	totalScore += CalculateMainStatScore(r, charWeight)
	fmt.Println("mainstats", r.Type, "=", totalScore)

	for i, sub := range r.SubAffix {
		val := sub.Value
		if sub.Percent {
			val = sub.Value * 100
		}

		weight := charWeight.SubstatWeights[sub.Type]
		effectiveValue := FindEffectiveStats(sub.Type)

		totalScore += (val / effectiveValue) * weight
		fmt.Println("sub", i, val, "/", effectiveValue, "*", weight, "=", math.Floor(((val/effectiveValue)*weight)*100)/100)
	}

	fmt.Println("final score ", totalScore, "*", (55 / partWeight), "=", totalScore*(55/partWeight))

	return totalScore * (55 / partWeight)
}

func GetSingleRelicRank(score float64) string {
	switch {
	case score >= 43.7:
		return "SSS"
	case score >= 37.8:
		return "SS"
	case score >= 33.5:
		return "S"
	case score >= 29.1:
		return "A"
	case score >= 23.3:
		return "B"
	case score >= 17.5:
		return "C"
	case score >= 11.6:
		return "D"
	default:
		return "N/A"
	}
}
