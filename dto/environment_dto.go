package dto

type EnvironmentResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddEnviroment struct {
	TeamID string `json:"team_id" binding:"required"`
	Key    string `json:"key" binding:"required"`
	Value  string `json:"value" binding:"required"`
}
