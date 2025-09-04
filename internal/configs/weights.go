package configs

import (
	"encoding/json"
	"os"

	"hsr-profile-tracker/internal/model"
)

var CharacterWeights map[string]model.CharacterWeights
var EffectiveStats *model.EffectiveStats

func LoadCharacterWeights(path string) (map[string]model.CharacterWeights, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chars []model.CharacterWeights
	if err := json.NewDecoder(file).Decode(&chars); err != nil {
		return nil, err
	}

	result := make(map[string]model.CharacterWeights)
	for _, char := range chars {
		result[char.Character] = char
	}

	return result, nil
}

func LoadEffectiveStats(path string) (*model.EffectiveStats, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result *model.EffectiveStats
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
