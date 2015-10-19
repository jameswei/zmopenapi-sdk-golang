package request

type BaseRequest struct {
	AppId      string
	MerchantId string
}

type BaseRequestInterface interface {
	GetParams() map[string]string
	GetAppId() string
}
