package model

type StatWeight struct {
	CoefficientStat map[string]float64 `json:"CoefficientStat"`
	BaseStat        int                `json:"BaseStat"`
}

type MainStats struct {
	Body   []string `json:"Body"`
	Feet   []string `json:"Feet"`
	Sphere []string `json:"Sphere"`
	Rope   []string `json:"Rope"`
}

type CharacterWeights struct {
	Character       string             `json:"character"`
	Role            string             `json:"role"`
	SubstatWeights  map[string]float64 `json:"substat_weights"`
	RelicPartWeight map[string]float64 `json:"relic_part_weight"`
	MainStats       MainStats          `json:"main_stats"`
	UsableSets      []string           `json:"usable_sets"`
}
