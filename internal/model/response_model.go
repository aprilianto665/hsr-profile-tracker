package model

type APIProfileResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    ProfileSummary `json:"data"`
}

type CheckProfileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}
