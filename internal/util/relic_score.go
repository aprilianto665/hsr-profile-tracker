package util

import (
	"fmt"
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/model"
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

func FindBaseStat(stat string) int {
	return configs.StatWeights.BaseStat[stat]
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
	// charWeight := FindCharacterWeights(player, char)

	// var totalScore float64

	// totalScore += CalculateMainStatScore(r, charWeight)
	// fmt.Println("mainstats", r.Type, "=", totalScore)

	for _, sub := range r.SubAffix {

		// val := sub.Value
		if sub.Type != "SpeedDelta" && !sub.Percent {
			fmt.Println("Stat Flat", sub.Type, FindBaseStat(sub.Type))
		}

		fmt.Println("Coefficient", sub.Type, FindStatCoefficient(sub.Type))

		// 	weight := charWeight.SubstatWeights[sub.Type]
		// 	effectiveValue := FindEffectiveStats(sub.Type)

		// 	totalScore += (val / effectiveValue) * weight
		// 	fmt.Println("sub", i, val, "/", effectiveValue, "*", weight, "=", math.Floor(((val/effectiveValue)*weight)*100)/100)
	}

	return 0
}

func GetSingleRelicRank(score float64) string {
	switch {
	case score >= 40:
		return "SSS"
	case score >= 35:
		return "SS"
	case score >= 30:
		return "S"
	case score >= 20:
		return "A"
	case score >= 15:
		return "B"
	case score >= 10:
		return "C"
	case score >= 0:
		return "D"
	default:
		return "N/A"
	}
}
