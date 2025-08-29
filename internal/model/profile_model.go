package model

type LightConeAttribute struct {
    Name    string  `json:"name"`
    Icon    string  `json:"icon"`
    Value   float64 `json:"value"`
    Percent bool    `json:"percent"` 
}

type LightCone struct {
    Name       string `json:"name"`
    Rarity     int    `json:"rarity"`
    Rank       int    `json:"rank"`
    Level      int    `json:"level"`
    Icon       string `json:"icon"`
    Attributes []LightConeAttribute `json:"attributes"`
}

type NameIcon struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Character struct {
	Name      string     `json:"name"`
	Portrait  string     `json:"portrait"`
    Rarity    int        `json:"rarity"`
	Rank      int        `json:"rank"`
	Level     int        `json:"level"`
    Path      *NameIcon  `json:"path"`
    Element   *NameIcon  `json:"element"`
    LightCone *LightCone `json:"light_cone"`
}


type SpaceInfo struct {
	UniverseLevel    int         `json:"universe_level"`
	AvatarCount      int         `json:"avatar_count"`
	LightConeCount   int         `json:"light_cone_count"`
	RelicCount       int         `json:"relic_count"`
	AchievementCount int         `json:"achievement_count"`
	BookCount        int         `json:"book_count"`
	MusicCount       int         `json:"music_count"`
}

type Player struct {
	UID      	string 			`json:"uid"`
	Nickname 	string 			`json:"nickname"`
	Level    	int    			`json:"level"`
	WorldLevel 	int  			`json:"world_level"`
	Friends 	int 			`json:"friend_count"`
	Signature 	string 			`json:"signature"`
	Avatar		*NameIcon		`json:"avatar"`
	SpaceInfo	*SpaceInfo		`json:"space_info"`
}

type RawData struct {
	Player	    Player          `json:"player"`
	Characters	[]Character	    `json:"characters"`
}