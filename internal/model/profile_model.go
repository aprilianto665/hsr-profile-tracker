package model

type RelicSet struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Num  int    `json:"num"`
}

type Attribute struct {
	Name    string  `json:"name"`
	Icon    string  `json:"icon"`
	Value   float64 `json:"value"`
	Percent bool    `json:"percent"`
}

type Relic struct {
	Name      string      `json:"name"`
	Type      int         `json:"type"`
	Rarity    int         `json:"rarity"`
	Level     int         `json:"level"`
	Icon      string      `json:"icon"`
	MainAffix *Attribute  `json:"main_affix"`
	SubAffix  []Attribute `json:"sub_affix"`
}

type LightCone struct {
	Name       string      `json:"name"`
	Rarity     int         `json:"rarity"`
	Rank       int         `json:"rank"`
	Level      int         `json:"level"`
	Icon       string      `json:"icon"`
	Attributes []Attribute `json:"attributes"`
}

type NameIcon struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Character struct {
	Name       string      `json:"name"`
	Portrait   string      `json:"portrait"`
	Rarity     int         `json:"rarity"`
	Rank       int         `json:"rank"`
	Level      int         `json:"level"`
	Path       *NameIcon   `json:"path"`
	Element    *NameIcon   `json:"element"`
	LightCone  *LightCone  `json:"light_cone"`
	Relics     []Relic     `json:"relics"`
	RelicSets  []RelicSet  `json:"relic_sets"`
	Attributes []Attribute `json:"attributes"`
	Additions  []Attribute `json:"additions"`
}

type SpaceInfo struct {
	UniverseLevel    int `json:"universe_level"`
	AvatarCount      int `json:"avatar_count"`
	LightConeCount   int `json:"light_cone_count"`
	RelicCount       int `json:"relic_count"`
	AchievementCount int `json:"achievement_count"`
	BookCount        int `json:"book_count"`
	MusicCount       int `json:"music_count"`
}

type Player struct {
	UID        string     `json:"uid"`
	Nickname   string     `json:"nickname"`
	Level      int        `json:"level"`
	WorldLevel int        `json:"world_level"`
	Friends    int        `json:"friend_count"`
	Signature  string     `json:"signature"`
	Avatar     *NameIcon  `json:"avatar"`
	SpaceInfo  *SpaceInfo `json:"space_info"`
}

type RawData struct {
	Player     Player      `json:"player"`
	Characters []Character `json:"characters"`
}

type RelicScore struct {
	Rank      string  `json:"rank"`
	BaseScore float64 `json:"base_score"`
	SetBonus  float64 `json:"set_bonus"`
}

type AttributeSummary struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Value string `json:"value"`
}

type RelicSummary struct {
	Name      string             `json:"name"`
	Type      int                `json:"type"`
	Rarity    int                `json:"rarity"`
	Level     int                `json:"level"`
	Icon      string             `json:"icon"`
	MainAffix *AttributeSummary  `json:"main_affix"`
	SubAffix  []AttributeSummary `json:"sub_affix"`
}

type LightConeSummary struct {
	Name       string             `json:"name"`
	Rarity     int                `json:"rarity"`
	Rank       int                `json:"rank"`
	Level      int                `json:"level"`
	Icon       string             `json:"icon"`
	Attributes []AttributeSummary `json:"attributes"`
}

type CharacterSummary struct {
	Name       string             `json:"name"`
	Portrait   string             `json:"portrait"`
	Rarity     int                `json:"rarity"`
	Rank       int                `json:"rank"`
	Level      int                `json:"level"`
	Path       *NameIcon          `json:"path"`
	Element    *NameIcon          `json:"element"`
	LightCone  *LightConeSummary  `json:"light_cone"`
	Relics     []RelicSummary     `json:"relics"`
	RelicSets  []RelicSet         `json:"relic_sets"`
	FinalStats []AttributeSummary `json:"final_stats"`
	RelicScore *RelicScore        `json:"relic_score"`
}

type ProfileSummary struct {
	Player     Player             `json:"player"`
	Characters []CharacterSummary `json:"characters"`
}
