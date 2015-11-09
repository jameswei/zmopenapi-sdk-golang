package request

import (
	"encoding/json"
	"zmopenapi-sdk/util"
)

const (
	MethodNameAuthorization = "zhima.auth.info.authorize"
)

type AuthorizationRequest struct {
	ProductID     string
	IdentityType  string
	IdentityParam map[string]string
	Channel       string
	Platform      string
	Sences        string
	BizParams     map[string]string
	ExtParams     map[string]string
}

func (r *AuthorizationRequest) GetParams() map[string]string {
	params := make(map[string]string, 10)
	params[util.ParamNameIdentityType] = r.IdentityType
	if r.IdentityParam != nil && len(r.IdentityParam) > 0 {
		json, err := json.Marshal(r.IdentityParam)
		if err == nil {
			params[util.ParamNameIdentityParam] = string(json)
		}
	}
	if r.BizParams != nil && len(r.BizParams) > 0 {
		json, err := json.Marshal(r.BizParams)
		if err == nil {
			params[util.ParamNameBizParams] = string(json)
		}
	}
	return params
}

func (r *AuthorizationRequest) GetAPIMethodName() string {
	return MethodNameAuthorization
}
