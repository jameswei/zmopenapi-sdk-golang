package request

import (
	"crypto"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"zmopenapi-sdk/util"
)

const (
	EmptyString = ""
	SymbolAnd   = "&"
	SymbolEqual = "="
	BodyType    = "application/x-www-form-urlencoded;charset="
)

var EncryptionFailed error = errors.New("encryption failed")

func BuildQuery(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return EmptyString
	}
	array := make([]string, 0, len(params))
	for key, value := range params {
		if key == EmptyString || value == EmptyString {
			continue
		}
		array = append(array, key+SymbolEqual+value)
	}
	return string(strings.Join(array, SymbolAnd))
}

func Encrypt(query string) (string, error) {
	return encryptQuery(query)
}

func encryptQuery(query string) (string, error) {
	encryptedQuery := util.EncryptBase64(util.EncryptRSA([]byte(query)))
	if encryptedQuery == EmptyString {
		return EmptyString, EncryptionFailed
	}
	return url.QueryEscape(encryptedQuery), nil
}

func Sign(query string) (string, error) {
	return signQuery(query)
}

func signQuery(query string) (string, error) {
	signature := util.EncryptBase64(util.SignRSA([]byte(query), crypto.SHA1))
	if signature == EmptyString {
		return EmptyString, EncryptionFailed
	}
	return url.QueryEscape(signature), nil
}

func BuildURL(r *CommonRequest) (string, error) {
	query := BuildQuery(r.GetParams())
	return encrytQuery(query, r.AppID)
}

func encrytQuery(query string, appId string) (string, error) {
	encrypted := make([]string, 0)
	encrypted = append(encrypted, util.ParamNameParams, SymbolEqual)
	encryptedParams := util.EncryptBase64(util.EncryptRSA([]byte(query)))
	if encryptedParams == EmptyString {
		return EmptyString, EncryptionFailed
	}
	encrypted = append(encrypted, url.QueryEscape(encryptedParams))
	encrypted = append(encrypted, SymbolAnd, util.ParamNameAppId, SymbolEqual, appId)
	encrypted = append(encrypted, SymbolAnd, util.ParamNameCharset, SymbolEqual, util.DefaultCharset)
	encrypted = append(encrypted, SymbolAnd, util.ParamNameSign, SymbolEqual)
	signature := util.EncryptBase64(util.SignRSA([]byte(query), crypto.MD5))
	if signature == EmptyString {
		return EmptyString, EncryptionFailed
	}
	encrypted = append(encrypted, url.QueryEscape(signature))
	encryptedQueryString := strings.Join(encrypted, EmptyString)
	return encryptedQueryString, nil
}

func HttpPost(client *http.Client, url string, content string, charset string) (string, error) {
	resp, err := client.Post(url, BodyType+charset, strings.NewReader(content))
	if err != nil {
		return EmptyString, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EmptyString, err
	}
	return string(body), nil
}
