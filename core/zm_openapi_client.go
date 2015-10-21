package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"zmopenapi-sdk/request"
	"zmopenapi-sdk/response"
	"zmopenapi-sdk/util"
)

var ErrorNetwork error = errors.New("network error!")
var ErrorInvalidData error = errors.New("invalid response data!")
var ErrorForgedData error = errors.New("forged response data!")

type ZMOpenApiClient interface {
	ExecuteCommonRequest(req *request.CommonRequest) (*response.CommonResponse, error)
	ExecuteFeedbackRequest(req *request.FeedbackRequest) (*response.FeedbackResponse, error)
	ExecuteCheckAuthRequest(req *request.CheckAuthRequest) (*response.CheckAuthResponse, error)
	ExeucteAuthEngineRequest(req *request.AuthEngineRequest) (*response.AuthEngineResponse, error)
	BuildAuthUrl(req *request.AuthEngineRequest) (string, error)
	BuildAgreementUrl(req *request.AgreementRequest) (string, error)
}

type DefaultZMOpenApiClient struct {
	PublicKey  string
	PrivateKey string
	Host       string
	charset    string
	AppId      string
	MerchantId string
	c          *http.Client
}

//NewZMOpenApiClient return a new ZMOpenApiClient
func NewZMOpenApiClient(publicKey, privateKey, appId, merchantId string) *ZMOpenApiClient {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 3),
	}
	return &DefaultZMOpenApiClient{
		c:          client,
		charset:    util.DefaultCharset,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		AppId:      appId,
		MerchantId: merchantId,
	}
}

func (client *DefaultZMOpenApiClient) ExecuteCommonRequest(req *request.CommonRequest) (*response.CommonResponse, error) {
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	if err != nil {
		return nil, err
	}
	response, err := request.HttpPost(client.c, client.Host+util.CommonURI, content, client.charset)
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
		decrypted := util.DecryptRSA(data)
		if util.VerifySignature(decrypted, commonResponse.SignResponse) {
			commonResponse.Result = string(decrypted)
		} else {
			return nil, ErrorForgedData
		}
	}
	return commonResponse, nil
}

func (client *DefaultZMOpenApiClient) ExecuteFeedbackRequest(req *request.FeedbackRequest) (*response.FeedbackResponse, error) {
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	if err != nil {
		return nil, err
	}
	response, err := request.HttpPost(client.c, client.Host+util.FeedbackURI, content, client.charset)
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
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	if err != nil {
		return nil, err
	}
	response, err := request.HttpPost(client.c, client.Host+util.CheckAuthURI, content, client.charset)
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
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	if err != nil {
		return nil, err
	}
	response, err := request.HttpPost(client.c, client.Host+util.AuthEngineURI, content, client.charset)
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

func (client *DefaultZMOpenApiClient) BuildAuthUrl(req *request.AuthEngineRequest) (string, error) {
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	return client.Host + util.AuthEngineURI + "?" + content, err
}

func (client *DefaultZMOpenApiClient) BuildAgreementUrl(req *request.AgreementRequest) (string, error) {
	req.AppId = client.AppId
	req.MerchantId = client.MerchantId
	content, err := request.BuildUrl(req)
	return client.Host + util.AgreementURI + "?" + content, err
}
