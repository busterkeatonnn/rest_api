package dto

type InfoResponse struct {
	PayloadResponse   interface{} `json:"payload_response"`
	TextResponse      string      `json:"text_response"`
	IsResponseSuccess bool        `json:"is_response_success"`
}
