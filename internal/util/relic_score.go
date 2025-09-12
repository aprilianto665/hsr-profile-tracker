package util

import (
	"fmt"
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/model"
)

func FindCharacterWeights(char model.Character) model.CharacterWeights {
	result := configs.CharacterWeights[char.Id]

	return result
}

func FindStatCoefficient(stat string) float64 {
	return configs.StatWeights.CoefficientStat[stat]
}

func FindBaseStat(stat string) float64 {
	return configs.StatWeights.BaseStat[stat]
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CheckRelicSuitableSet(r model.Relic, charWeight model.CharacterWeights) bool {
	return contains(charWeight.UsableSets, r.SetName)
}

func CalculateFinalStatScore(r model.Relic, charWeight model.CharacterWeights, score float64) float64 {
	penaltyMultiplier := 1.0

	isRelicSetSuitable := CheckRelicSuitableSet(r, charWeight)

	if !isRelicSetSuitable {
		penaltyMultiplier -= 0.1
	}

	slotMap := map[int]string{
		3: "Body",
		4: "Feet",
		5: "Sphere",
		6: "Rope",
	}

	fmt.Println(r.SetName)
	fmt.Println(r.MainAffix.Type)

	slotName := slotMap[r.Type]
	if slotName == "" {
		return FloorToDecimal(score*penaltyMultiplier, 1)
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
		return FloorToDecimal((score+5.832)*penaltyMultiplier, 1)
	} else {
		penaltyMultiplier -= 0.5
	}

	return FloorToDecimal(score*penaltyMultiplier, 1)
}

func CalculateRelicScoreValue(r model.Relic, char model.Character) float64 {
	charWeight := FindCharacterWeights(char)

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
	fmt.Println("total floor :", FloorToDecimal(totalScore, 1))
	totalScore = CalculateFinalStatScore(r, charWeight, FloorToDecimal(totalScore, 1))
	fmt.Println("total :", totalScore)
	return totalScore
}

func GetRelicRank(score float64) string {
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
	case score >= 5:
		return "C"
	case score > 0:
		return "D"
	default:
		return "N/A"
	}
}
