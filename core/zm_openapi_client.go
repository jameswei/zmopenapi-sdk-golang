package core

import (
	"crypto"
	"encoding/json"
	"errors"
	"linkedin/ginutil"
	"net/http"
	"time"
	"zmopenapi-sdk/request"
	"zmopenapi-sdk/response"
	"zmopenapi-sdk/util"
)

const (
	EmptyString = ""
)

var ErrorNetwork = errors.New("network error!")
var ErrorInvalidData = errors.New("invalid response data!")
var ErrorForgedData = errors.New("forged response data!")
var ErrorEncryptionFailed = errors.New("encryption failed!")

type ZMOpenAPIClient struct {
	Host       string
	AppID      string
	MerchantID string
	charset    string
	c          *http.Client
}

//NewZMOpenApiClient return a new ZMOpenApiClient
func NewZMOpenApiClient() *ZMOpenAPIClient {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 3),
	}
	ret := &ZMOpenAPIClient{
		Host:       util.DefaultHost,
		AppID:      util.AppID,
		MerchantID: util.MerchantID,
		charset:    util.DefaultCharset,
		c:          client,
	}
	if !ginutil.IsProduction() {
		ret.AppID = util.TestAppID
		ret.MerchantID = util.TestMerchantID
	}
	return ret
}

func (client *ZMOpenAPIClient) ExecuteCommonRequest(req *request.CommonRequest) (*response.CommonResponse, error) {
	req.AppID = client.AppID
	req.MerchantID = client.MerchantID
	content, err := request.BuildURL(req)
	if err != nil {
		return nil, err
	}
	resp, err := request.HttpPost(client.c, client.Host+util.CommonURI, content, client.charset)
	if err != nil {
		return nil, ErrorNetwork
	}
	var commonResponse response.CommonResponse
	err = json.Unmarshal([]byte(resp), &commonResponse)
	if err != nil {
		return nil, ErrorInvalidData
	}
	if commonResponse.Success {
		decrypted := util.DecryptRSA(util.DecryptBase64(commonResponse.Response))
		commonResponse.Result = string(decrypted)
		if util.VerifySignature(decrypted, commonResponse.SignResponse, crypto.MD5) {
			commonResponse.Result = string(decrypted)
		} else {
			return nil, ErrorForgedData
		}
	}
	return &commonResponse, nil
}

func (client *ZMOpenAPIClient) BuildAuthorizationInfo(req *request.AuthorizationRequest) (string, string, error) {
	query := request.BuildQuery(req.GetParams())
	if query == EmptyString || len(query) == 0 {
		return EmptyString, EmptyString, ErrorEncryptionFailed
	}
	encrypted, err := request.Encrypt(query)
	if err != nil {
		return EmptyString, EmptyString, ErrorEncryptionFailed
	}
	signature, err := request.Sign(query)
	if err != nil {
		return EmptyString, EmptyString, ErrorEncryptionFailed
	}
	return encrypted, signature, nil
}
