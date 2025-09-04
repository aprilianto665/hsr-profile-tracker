package configs

import (
	"encoding/json"
	"os"

	"hsr-profile-tracker/internal/model"
)

var CharacterWeights *model.CharacterWeights
var EffectiveStats *model.EffectiveStats

func LoadEffectiveStatstWeights(path string) (*model.EffectiveStats, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sw model.EffectiveStats
	if err := json.NewDecoder(f).Decode(&sw); err != nil {
		return nil, err
	}
	return &sw, nil
}

func LoadCharacterWeights(path string) (*model.CharacterWeights, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cw model.CharacterWeights
	if err := json.NewDecoder(f).Decode(&cw); err != nil {
		return nil, err
	}
	return &cw, nil
}
