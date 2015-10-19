package response

type CheckAuthResponse struct {
	Success      bool
	ErrorCode    string
	ErrorMessage string
	Response     string
}
