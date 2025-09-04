package configs

import (
	"encoding/json"
	"os"

	"hsr-profile-tracker/internal/model"
)

type WeightType interface {
	*model.EffectiveStats | *model.CharacterWeights
}

var CharacterWeights *model.CharacterWeights
var EffectiveStats *model.EffectiveStats

func LoadWeights[T WeightType](path string) (T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result T
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
