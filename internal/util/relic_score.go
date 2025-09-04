package util

import (
	"fmt"
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/model"
	"reflect"
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

func FindEffectiveStats(statType string) float64 {
	value := reflect.ValueOf(*configs.EffectiveStats)
	field := value.FieldByName(statType).Float()

	return field
}

func CalculateRelicScore(player model.Player, char model.Character) {
	charWeight := FindCharacterWeights(player, char)

	for _, r := range char.Relics {
		fmt.Println(r.Name)
		eff := FindEffectiveStats(r.MainAffix.Type)
		fmt.Println(r.MainAffix.Name, eff)

		for _, sub := range r.SubAffix {
			charStatWeight := charWeight.SubstatWeights[sub.Type]
			fmt.Println(sub.Type, charStatWeight)
		}
	}
}
