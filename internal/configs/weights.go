package configs

import (
	"encoding/json"
	"os"

	"hsr-profile-tracker/internal/model"
)

var CharacterWeights map[string]model.CharacterWeights
var StatWeights *model.StatWeight

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
		for _, id := range char.Id {
			result[id] = char
		}
	}

	return result, nil
}

func LoadStatWeights(path string) (*model.StatWeight, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result *model.StatWeight
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
