package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"zmopenapi-sdk/util"
)

func BuildUrl(r BaseRequestInterface) string {
	queryString := buildQuery(r.GetParams())
	return encrytQuery(queryString, r.GetAppId())
}

func buildUrl(r BaseRequestInterface, url string) string {
	queryString := buildQuery(r.GetParams())
	if strings.ContainsAny(url, "?") {
		return url + encrytQuery(queryString, r.GetAppId())
	} else {
		return url + "?" + encrytQuery(queryString, r.GetAppId())
	}
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

func encrytQuery(queryString string, appId string) string {
	encrypted := make([]string, 0)
	encrypted = append(encrypted, util.ParamNameParams, "=")
	encrypted = append(encrypted, url.QueryEscape(util.EncryptBase64(util.EncryptRSA([]byte(queryString)))))
	encrypted = append(encrypted, "&", util.ParamNameAppId, "=", appId)
	encrypted = append(encrypted, "&", util.ParamNameCharset, "=", util.DefaultCharset)
	encrypted = append(encrypted, "&", util.ParamNameSign, "=")
	encrypted = append(encrypted, url.QueryEscape(string(
		util.SignWithRSA(util.EncryptSHA([]byte(queryString))))))
	encryptedQueryString := strings.Join(encrypted, "")
	return encryptedQueryString
}

func HttpPost(url string, content string, charset string) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded;charset="+charset, strings.NewReader(content))
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
