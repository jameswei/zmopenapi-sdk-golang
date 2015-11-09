package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type CommonRequest struct {
	AppID      string
	MerchantID string
	CustomerID string
	ProductID  string
	Sences     string
	ExtParams  map[string]string
}

func (r *CommonRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameMerchantId] = r.MerchantID
	params[util.ParamNameCustomerId] = r.CustomerID
	params[util.ParamNameProductId] = r.ProductID
	params[util.ParamNameSences] = r.Sences
	if r.ExtParams != nil && len(r.ExtParams) > 0 {
		json, err := json.Marshal(r.ExtParams)
		if err == nil {
			params[util.ParamNameExtParams] = string(json)
		}
	}
	return params
}
