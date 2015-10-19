package response

type AuthEngineResponse struct {
	Success      bool
	ErrorCode    string
	ErrorMessage string
	Response     string
}
