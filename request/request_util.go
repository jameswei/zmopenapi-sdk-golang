package request

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"zmopenapi-sdk/util"
)

var EncryptionFailed error = errors.New("encryption failed")

func BuildUrl(r BaseRequestInterface) (string, error) {
	queryString := buildQuery(r.GetParams())
	return encrytQuery(queryString, r.GetAppId())
}

func buildQuery(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	array := make([]string, 0, len(params))
	for key, value := range params {
		if key == "" || value == "" {
			continue
		}
		array = append(array, key+"="+value)
	}
	return string(strings.Join(array, "&"))
}

func encrytQuery(queryString string, appId string) (string, error) {
	encrypted := make([]string, 0)
	encrypted = append(encrypted, util.ParamNameParams, "=")
	encryptedParams := util.EncryptBase64(util.EncryptRSA([]byte(queryString)))
	if encryptedParams == "" {
		return "", EncryptionFailed
	}
	encrypted = append(encrypted, url.QueryEscape(encryptedParams))
	encrypted = append(encrypted, "&", util.ParamNameAppId, "=", appId)
	encrypted = append(encrypted, "&", util.ParamNameCharset, "=", util.DefaultCharset)
	encrypted = append(encrypted, "&", util.ParamNameSign, "=")
	signature := string(util.SignWithRSA([]byte(queryString)))
	if signature == "" {
		return "", EncryptionFailed
	}
	encrypted = append(encrypted, url.QueryEscape(signature))
	encryptedQueryString := strings.Join(encrypted, "")
	return encryptedQueryString, nil
}

func HttpPost(client *http.Client, url string, content string, charset string) (string, error) {
	resp, err := client.Post(url, "application/x-www-form-urlencoded;charset="+charset, strings.NewReader(content))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
