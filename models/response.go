package models

// Response represents the model of a response
type Response struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
