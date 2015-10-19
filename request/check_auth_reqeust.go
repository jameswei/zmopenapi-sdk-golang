package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type CheckAuthRequest struct {
	BaseRequest
	ProductID string
	AuthType  string
	AuthParam map[string]string
	Sences    string
}

func (r *CheckAuthRequest) GetAppId() string {
	return r.AppId
}

func (r *CheckAuthRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameMerchantId] = r.MerchantId
	params[util.ParamNameProductId] = r.ProductID
	params[util.ParamNameAuthType] = r.AuthType
	if r.AuthParam != nil && len(r.AuthParam) > 0 {
		b, err := json.Marshal(r.AuthParam)
		if err == nil {
			params[util.ParamNameAuthParam] = string(b)
		}
	}
	params[util.ParamNameSences] = r.Sences
	return params
}
