package domain

type ResponseDefault struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   bool        `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
