package model

type EffectiveStats struct {
	CriticalChanceBase        float64 `json:"CriticalChanceBase"`
	CriticalDamageBase        float64 `json:"CriticalDamageBase"`
	AttackAddedRatio          float64 `json:"AttackAddedRatio"`
	HPAddedRatio              float64 `json:"HPAddedRatio"`
	DefenceAddedRatio         float64 `json:"DefenceAddedRatio"`
	AttackDelta               float64 `json:"AttackDelta"`
	HPDelta                   float64 `json:"HPDelta"`
	DefenceDelta              float64 `json:"DefenceDelta"`
	SpeedDelta                float64 `json:"SpeedDelta"`
	StatusProbabilityBase     float64 `json:"StatusProbabilityBase"`
	StatusResistanceBase      float64 `json:"StatusResistanceBase"`
	BreakDamageAddedRatioBase float64 `json:"BreakDamageAddedRatioBase"`
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
