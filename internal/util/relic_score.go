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

func CalculateRelicScore(player model.Player, char model.Character) {
	data := FindCharacterWeights(player, char)
	fmt.Println(data.Character)
}
