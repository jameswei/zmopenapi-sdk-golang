package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

type AgreementRequest struct {
	BaseRequest
	AgreementType string
	State         string
	ExtParam      map[string]string
}

func (r *AgreementRequest) GetAppId() string {
	return r.AppId
}

func (r *AgreementRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameMerchantId] = r.MerchantId
	params[util.ParamNameAgreementType] = r.AgreementType
	params[util.ParamNameState] = r.State
	if r.ExtParam != nil && len(r.ExtParam) > 0 {
		b, err := json.Marshal(r.ExtParam)
		if err == nil {
			params[util.ParamNameExt] = string(b)
		}
	}
	return params
}
