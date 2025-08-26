package model

type AvatarInfo struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type MemoryInfo struct {
	Level          int  `json:"level"`
	ChaosID        *int `json:"chaos_id,omitempty"`
	ChaosLevel     *int `json:"chaos_level,omitempty"`
	ChaosStarCount *int `json:"chaos_star_count,omitempty"`
}

type SpaceInfo struct {
	MemoryData       *MemoryInfo `json:"memory_data,omitempty"`
	UniverseLevel    int         `json:"universe_level"`
	AvatarCount      int         `json:"avatar_count"`
	LightConeCount   int         `json:"light_cone_count"`
	RelicCount       int         `json:"relic_count"`
	AchievementCount int         `json:"achievement_count"`
	BookCount        int         `json:"book_count"`
	MusicCount       int         `json:"music_count"`
}

type Player struct {
	UID      	string 		`json:"uid"`
	Nickname 	string 		`json:"nickname"`
	Level    	int    		`json:"level"`
	WorldLevel 	int  		`json:"world_level"`
	Friends 	int 		`json:"friend_count"`
	Signature 	string 		`json:"signature"`
	Avatar		*AvatarInfo	`json:"avatar"`
	SpaceInfo	*SpaceInfo	`json:"space_info"`
}

type MihomoProfileResponse struct {
	Player	Player `json:"player"`
}

type APIProfileResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    Player `json:"data"`
}

type CheckProfileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}