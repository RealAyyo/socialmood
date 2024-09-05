package controllers

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Result  int         `json:"result,omitempty"`
}
