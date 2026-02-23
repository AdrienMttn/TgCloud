package types

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Id string `json:"id,omitempty"`
}