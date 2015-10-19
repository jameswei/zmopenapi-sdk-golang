package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type AuthEngineRequest struct {
	BaseRequest
	ProductID string
	AuthType  string
	AuthParam map[string]string
	Sences    string
	View      string
	PageType  string
	State     string
	Isv       string
	ExtParams map[string]string
}

func (r *AuthEngineRequest) GetAppId() string {
	return r.AppId
}

func (r *AuthEngineRequest) GetParams() map[string]string {
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
	params[util.ParamNameView] = r.View
	params[util.ParamNamePageType] = r.PageType
	params[util.ParamNameState] = r.State
	params[util.ParamNameISV] = r.Isv
	if r.ExtParams != nil && len(r.ExtParams) > 0 {
		b, err := json.Marshal(r.ExtParams)
		if err == nil {
			params[util.ParamNameExt] = string(b)
		}
	}
	return params
}
