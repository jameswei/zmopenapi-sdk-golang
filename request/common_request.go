package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type CommonRequest struct {
	BaseRequest
	CustomerID string
	ProductID  string
	Sences     string
	View       string
	State      string
	Isv        string
	ExtParams  map[string]string
}

func (r *CommonRequest) GetAppId() string {
	return r.AppId
}

func (r *CommonRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameMerchantId] = r.MerchantId
	params[util.ParamNameCustomerId] = r.CustomerID
	params[util.ParamNameProductId] = r.ProductID
	params[util.ParamNameSences] = r.Sences
	params[util.ParamNameView] = r.View
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
