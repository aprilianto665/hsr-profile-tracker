package model

type StatWeight struct {
	CoefficientStat map[string]float64 `json:"coefficient_stat"`
	BaseStat        map[string]int     `json:"base_stat"`
}

type MainStats struct {
	Body   []string `json:"Body"`
	Feet   []string `json:"Feet"`
	Sphere []string `json:"Sphere"`
	Rope   []string `json:"Rope"`
}

type CharacterWeights struct {
	Character      string             `json:"character"`
	SubstatWeights map[string]float64 `json:"substat_weights"`
	MainStats      MainStats          `json:"main_stats"`
	UsableSets     []string           `json:"usable_sets"`
}
