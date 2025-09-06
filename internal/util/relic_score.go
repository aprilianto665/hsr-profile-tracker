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

func FindBaseStat(stat string) float64 {
	return configs.StatWeights.BaseStat[stat]
}

func CalculateMainStatScore(r model.Relic, charWeight model.CharacterWeights, score float64) float64 {
	slotMap := map[int]string{
		3: "Body",
		4: "Feet",
		5: "Sphere",
		6: "Rope",
	}

	slotName := slotMap[r.Type]
	if slotName == "" {
		return score
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
		return score + 5.832
	}

	return score * 0.5
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

	var totalScore float64

	fmt.Println("Relic Slot", r.Type)

	for _, sub := range r.SubAffix {

		if sub.Type != "SpeedDelta" && !sub.Percent {

			totalScore += ((sub.Value / FindBaseStat(sub.Type)) * 100) * FindStatCoefficient(sub.Type) * charWeight.SubstatWeights[sub.Type]
			fmt.Println("Stat Flat", sub.Type, sub.Value, "/", FindBaseStat(sub.Type), "* 100 *", FindStatCoefficient(sub.Type), "*", charWeight.SubstatWeights[sub.Type], "=", ((sub.Value/FindBaseStat(sub.Type))*100)*FindStatCoefficient(sub.Type)*charWeight.SubstatWeights[sub.Type])
		} else {

			val := sub.Value

			if sub.Type != "SpeedDelta" {
				val *= 100
			}

			totalScore += val * FindStatCoefficient(sub.Type) * charWeight.SubstatWeights[sub.Type]
			fmt.Println("Stat Percent", sub.Type, val, "*", FindStatCoefficient(sub.Type), "*", charWeight.SubstatWeights[sub.Type], "=", val*FindStatCoefficient(sub.Type)*charWeight.SubstatWeights[sub.Type])
		}
	}

	totalScore = CalculateMainStatScore(r, charWeight, totalScore)
	fmt.Println("total :", totalScore)
	return totalScore
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
