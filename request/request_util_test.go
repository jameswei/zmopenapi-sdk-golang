package request

import (
	"fmt"
	"github.com/qiniu/api.v6/url"
	"testing"
	"zmopenapi-sdk/util"
)

func TestBuildQuery(t *testing.T) {
	params := make(map[string]string)
	params["ka"] = "va"
	params["kb"] = "vb"
	params["kc"] = "vc"
	queryString := buildQuery(params)
	if queryString != "ka=va&kb=vb&kc=vc" {
		t.Error("BuildQuery failed")
	}
}

func TestBuildUrl(t *testing.T) {
	host := "https://zmopenapi.zmxy.com.cn/"
	authEngineRequest := &AuthEngineRequest{}
	authEngineRequest.AppId = "1000080"
	authEngineRequest.MerchantId = "268820000049398445673"
	authEngineRequest.AuthType = "0"
	authEngineRequest.View = "redirect"
	authEngineRequest.Sences = "chitu"
	authEngineRequest.PageType = "pc"
	authParams := make(map[string]string, 2)
	authParams["certNo"] = "650104198505260712"
	authParams["name"] = "魏佳"
	authEngineRequest.AuthParam = authParams
	url, err := BuildUrl(authEngineRequest)
	if err != nil {
		t.Error(err)
	}
	builtUrl := host + "authorize.do" + "?" + url
	fmt.Println(builtUrl)

	agreementRequest := &AgreementRequest{}
	agreementRequest.AppId = "1000080"
	agreementRequest.MerchantId = "268820000049398445673"
	agreementRequest.AgreementType = "3"
	url, err = BuildUrl(agreementRequest)
	if err != nil {
		t.Error(err)
	}
	builtUrl = host + "agreement.do" + "?" + url
	fmt.Println(builtUrl)

	checkAuthRequest := &CheckAuthRequest{}
	checkAuthRequest.AppId = "1000080"
	checkAuthRequest.MerchantId = "268820000049398445673"
	checkAuthRequest.AuthType = "0"
	checkAuthRequest.Sences = "chitu"
	authParams = make(map[string]string, 2)
	authParams["certNo"] = "650104198505260712"
	authParams["name"] = "魏佳"
	checkAuthRequest.AuthParam = authParams
	url, err = BuildUrl(checkAuthRequest)
	if err != nil {
		t.Error(err)
	}
	builtUrl = host + "authorized.do" + "?" + url
	fmt.Println(builtUrl)

	feedbackRequest := &FeedbackRequest{}
	feedbackRequest.AppId = "1000080"
	feedbackRequest.MerchantId = "268820000049398445673"
	feedbackRequest.Data = "test-data"
	feedbackRequest.Identity = "test-identity"
	feedbackRequest.TypeID = "test-typeId"
	url, err = BuildUrl(feedbackRequest)
	if err != nil {
		t.Error(err)
	}
	builtUrl = host + "feedback.do" + "?" + url
	fmt.Println(builtUrl)
}

func TestDecrypt(t *testing.T) {
	params := "Srf9T7O1DGXw6E5m4XqMYEXurRfWnIpAULup%2B9JbZiKiWYGfhphYHddToYD6GrwfjhlIkO1YcY2gK%2FE6mTj4MwV0%2BmccwaJgxFa4PT3hSwcAjXtHqiezT7Tw5LxCkEnEwqLfKuleaIpoiLUR3s2EBrZ3F%2F7AP%2FAyKbpUEphNRNA%3D"
	unescaped, err := url.Unescape(params)
	if err != nil {
		t.Error(err)
	}
	data := util.DecryptBase64(unescaped)
	decrypted := string(util.DecryptRSA(data))
	fmt.Println("decrypted:", decrypted)
}

func TestSignature(t *testing.T) {
	params := "Srf9T7O1DGXw6E5m4XqMYEXurRfWnIpAULup%2B9JbZiKiWYGfhphYHddToYD6GrwfjhlIkO1YcY2gK%2FE6mTj4MwV0%2BmccwaJgxFa4PT3hSwcAjXtHqiezT7Tw5LxCkEnEwqLfKuleaIpoiLUR3s2EBrZ3F%2F7AP%2FAyKbpUEphNRNA%3D"
	unescaped, err := url.Unescape(params)
	if err != nil {
		t.Error(err)
	}
	signature := "UlfDNbAzhUCvCsfi7hgTGM5AQ8kFYRZ84kT3Cm1Ofem41J1cf4sH8sbrSOwu%2F9Osv%2BoDumAA63yS%2B7EpEuMeUGs%2Fy8XvIGFOKaZb2MPsbDULIU%2B2vtjB7jh2N9fIzz24CvoH2e136K8oqA%2FtEhPN7Xzx%2FgAVadwL8ZVE6vOyhPU%3D"
	unescapedSignature, err := url.Unescape(signature)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("escaped sign:", unescapedSignature)
	data := util.DecryptBase64(unescaped)
	decrypted := util.DecryptRSA(data)
	verified := util.VerifySignature(decrypted, unescapedSignature)
	if !verified {
		t.Error(verified)
	}
}
