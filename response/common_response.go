package response

type CommonResponse struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Response     string `json:"response,omitempty"`
	SignResponse string `json:"signResponse,omitempty"`
	Result       string `json:"result,omitempty"`
}

type ZhiMaScore struct {
	Score int32 `json:"score"`
}

type ZhiMaScoreResult struct {
	Content   *ZhiMaScore `json:"content"`
	Zmproduct string      `json:"zmproduct"`
}
