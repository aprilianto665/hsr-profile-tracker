package model

type APIProfileResponse struct {
    Status  string  `json:"status"`
    Message string  `json:"message"`
    Data    RawData `json:"data"`
}

type CheckProfileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Exists  bool   `json:"exists"`
}