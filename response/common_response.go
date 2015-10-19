package response

type CommonResponse struct {
	Success      bool
	ErrorCode    string
	ErrorMessage string
	Response     string
	SignResponse string
	Result       string
}
