package response

type AuthorizeResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Response     string `json:"response,omitempty"`
}
