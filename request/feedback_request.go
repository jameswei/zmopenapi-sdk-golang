package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type FeedbackRequest struct {
	BaseRequest
	TypeID    string
	Identity  string
	Data      string
	ExtParams map[string]string
}

func (r *FeedbackRequest) GetAppId() string {
	return r.AppId
}

func (r *FeedbackRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameMerchantId] = r.MerchantId
	params[util.ParamNameTypeId] = r.TypeID
	params[util.ParamNameIdentity] = r.Identity
	params[util.ParamNameData] = r.Data
	if r.ExtParams != nil && len(r.ExtParams) > 0 {
		b, err := json.Marshal(r.ExtParams)
		if err == nil {
			params[util.ParamNameExt] = string(b)
		}
	}
	return params
}
