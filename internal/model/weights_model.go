package model

type EffectiveStats struct {
	CriticalChanceBase     float64 `json:"CriticalChanceBase"`
	CriticalDamageBase     float64 `json:"CriticalDamageBase"`
	AttackAddedRatio       float64 `json:"AttackAddedRatio"`
	HPAddedRatio           float64 `json:"HPAddedRatio"`
	DefenceAddedRatio      float64 `json:"DefenceAddedRatio"`
	AttackDelta            float64 `json:"AttackDelta"`
	HPDelta                float64 `json:"HPDelta"`
	DefenceDelta           float64 `json:"DefenceDelta"`
	SpeedDelta             float64 `json:"SpeedDelta"`
	StatusProbabilityBase  float64 `json:"StatusProbabilityBase"`
	StatusResistanceBase   float64 `json:"StatusResistanceBase"`
	BreakDamageAddedRatio  float64 `json:"BreakDamageAddedRatio"`
	EnergyRegenerationRate float64 `json:"EnergyRegenerationRate"`
	FireAddedRatio         float64 `json:"FireAddedRatio"`
	IceAddedRatio          float64 `json:"IceAddedRatio"`
	ImaginaryAddedRatio    float64 `json:"ImaginaryAddedRatio"`
	LightningAddedRatio    float64 `json:"LightningAddedRatio"`
	PhysicalAddedRatio     float64 `json:"PhysicalAddedRatio"`
	QuantumAddedRatio      float64 `json:"QuantumAddedRatio"`
	WindAddedRatio         float64 `json:"WindAddedRatio"`
}

type MainStats struct {
	Body   []string `json:"Body"`
	Feet   []string `json:"Feet"`
	Sphere []string `json:"Sphere"`
	Rope   []string `json:"Rope"`
}

type CharacterWeights struct {
	Character      string         `json:"character"`
	Role           string         `json:"role"`
	EffectiveStats EffectiveStats `json:"effective_stats"`
	MainStats      MainStats      `json:"main_stats"`
	UsableSets     []string       `json:"usable_sets"`
}
