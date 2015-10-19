package core

import (
	"encoding/json"
	"errors"
	"zmopenapi-sdk/request"
	"zmopenapi-sdk/response"
	"zmopenapi-sdk/util"
)

const (
	ErrorNetwork     = errors.New("network error!")
	ErrorInvalidData = errors.New("invalid response data!")
	ErrorForgedData  = errors.New("forged response data!")
)

type ZMOpenApiClient interface {
	ExecuteCommonRequest(req *request.CommonRequest) (*response.CommonResponse, error)
	ExecuteFeedbackRequest(req *request.FeedbackRequest) (*response.FeedbackResponse, error)
	ExecuteCheckAuthRequest(req *request.CheckAuthRequest) (*response.CheckAuthResponse, error)
	ExeucteAuthEngineRequest(req *request.AuthEngineRequest) (*response.AuthEngineResponse, error)
}

type DefaultZMOpenApiClient struct {
	PublicKey         string
	PrivateKey        string
	Charset           string
	ConnectionTimeout uint32
	ReadTimeout       uint32
	Host              string
}

func (client *DefaultZMOpenApiClient) ExecuteCommonRequest(req *request.CommonRequest) (*response.CommonResponse, error) {
	content := request.BuildUrl(req)
	response, err := request.HttpPost(client.Host+util.CommonURI, content, client.Charset)
	if err != nil {
		return nil, ErrorNetwork
	}
	var commonResponse *response.CommonResponse
	err = json.Unmarshal(response, commonResponse)
	if err != nil {
		return nil, ErrorInvalidData
	}
	if commonResponse.Success {
		encryptedResponse := commonResponse.Response
		data := util.DecryptBase64(encryptedResponse)
		result := util.DecryptRSA(data)
		generated := util.EncryptSHA(result)
		if generated == util.VerifySignature(generated, commonResponse.SignResponse) {
			commonResponse.Result = result
		} else {
			return nil, ErrorForgedData
		}
	}
	return commonResponse, nil
}

func (client *DefaultZMOpenApiClient) ExecuteFeedbackRequest(req *request.FeedbackRequest) (*response.FeedbackResponse, error) {
	content := request.BuildUrl(req)
	response, err := request.HttpPost(client.Host+util.FeedbackURI, content, client.Charset)
	if err != nil {
		return nil, ErrorNetwork
	}
	var feedbackResponse *response.FeedbackResponse
	err = json.Unmarshal(response, feedbackResponse)
	if err != nil {
		return nil, ErrorInvalidData
	}
	return feedbackResponse, nil
}

func (client *DefaultZMOpenApiClient) ExecuteCheckAuthRequest(req *request.CheckAuthRequest) (*response.CheckAuthResponse, error) {
	content := request.BuildUrl(req)
	response, err := request.HttpPost(client.Host+util.CheckAuthURI, content, client.Charset)
	if err != nil {
		return nil, ErrorNetwork
	}
	var checkAuthResponse *response.CheckAuthResponse
	err = json.Unmarshal(response, checkAuthResponse)
	if err != nil {
		return nil, ErrorInvalidData
	}
	return checkAuthResponse, nil
}

func (client *DefaultZMOpenApiClient) ExecuteAuthEngineRequest(req *request.AuthEngineRequest) (*response.AuthEngineResponse, error) {
	content := request.BuildUrl(req)
	response, err := request.HttpPost(client.Host+util.AuthEngineURI, content, client.Charset)
	if err != nil {
		return nil, ErrorNetwork
	}
	var authEngineResponse *response.AuthEngineResponse
	err = json.Unmarshal(response, authEngineResponse)
	if err != nil {
		return nil, ErrorInvalidData
	}
	return authEngineResponse, nil
}
