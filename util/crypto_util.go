package util

import (
	"crypto"
	md5 "crypto/md5"
	rand "crypto/rand"
	rsa "crypto/rsa"
	sha1 "crypto/sha1"
	x509 "crypto/x509"
	base64 "encoding/base64"
	pem "encoding/pem"
)

const (
	AlgorithmSHA = "SHA"
	AlgorithmMD5 = "MD5"
	AlgorithmRSA = "RSA"
	PrivateKey   = `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAMbkdtxCjpSZju3x
+CEal6m+PojEg9iqdqiv6MTd1UwDcIIXWeZcXGI9PeLYFy9V/qTtqezdrVWLSO5m
PZhb6hh5x2VqoDvAXx6MnOOwhlqmSJxkxBhUK/cmM0k7Uzcd5k6cMmQOt6WsF+Sc
9+g0iZOl3nJonrGTPlhHgvmKfKlTAgMBAAECgYB9tsqqTidxJd6B5/++bOCQGf/M
0unDeXunBAlM5ip78XCbyca5JIgAUFVdnNiKwwBBnzdY0IVPHMrsZRNpyi8cSyxi
x7+lw5COK8Tg5XG99vadQotVw1Sx8tT2xFcwtf+PZBrj3jtiOWeO/Qrk3nVF4Fup
FXb69Ho4nYFiDZmYgQJBAPSvwUpzpgonauSUr5thUY4U13QdZVvL8tE6KXXPh+NP
16ooPI5dnYQRJMXA7Ejalj/tCcNE5991bmHYeH4JNu0CQQDQFqo7DfKKRp82ATQe
IF36gozq3O7NAETBTYIZed05sjdxIrni1zYBzFu+7Zp3VpB/JdsK/7SrP8kFcY1Q
5xk/AkEAimF7p2eQV93DDlMonW+EeB5BW2HkmO3W/Y0vNXmRGHVnOsxWsw0usCoh
6dUZzHoSz0R3aP/nQvFe4+dQ/baoYQJAMVt+610Gj6fqscudShwRTo9Sz46yEewj
Ytp4poRSZhIQtoQvJVA43jpT9Li1L+NXiOE40KYQe6I4k1L42xzFzwJBAKYUJx+b
gaNO/6oPL1fJogHWk/pYQMxmexKV0xEO6t3p/7tE1kshFx9Lx1ydv5Q7yqt66vcC
YuGheYjYNrO3Ec8=
-----END PRIVATE KEY-----`
	PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCScu8DUo7vNEy7oySv49zn5WiE
JIjayPcKqw63uEGa0Y1Kpr4xWstURwK9eFAP7wC0wtG833HLuxWnQOqm4at3MrPj
WsVUXnWl9YhIKjWjqJwDFtZgemczLM3+1waaINxZCby17K7+qkljXoDc/s9hOZwo
6FaDEq6ByENV2warmQIDAQAB
-----END PUBLIC KEY-----`
)

//EncryptBase64 encrypt given []byte with Base64 algorithm
func EncryptBase64(data []byte) string {
	if data == nil {
		return ""
	}
	encrypted := base64.StdEncoding.EncodeToString(data)
	return encrypted
}

//DecryptBase64 decrypt given string with Base64 algorithm
func DecryptBase64(data string) []byte {
	if data == "" {
		return nil
	}
	decrypted, _ := base64.StdEncoding.DecodeString(data)
	return decrypted
}

//EncryptMD5 encrypt given []byte with MD5 algorithm
func EncryptMD5(data []byte) []byte {
	if data == nil {
		return nil
	}
	encrypter := md5.New()
	encrypter.Write(data)
	return encrypter.Sum(nil)
}

//EncryptSHA encrypt given []byte with SHA algorithm
func EncryptSHA(data []byte) []byte {
	if data == nil {
		return nil
	}
	encypter := sha1.New()
	encypter.Write(data)
	return encypter.Sum(nil)
}

//EncryptRSA encrypt given data with RSA algorithm
func EncryptRSA(data []byte) []byte {
	if data == nil {
		return nil
	}
	block, _ := pem.Decode([]byte(PublicKey))
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	pub := pubInterface.(*rsa.PublicKey)
	encrypted := make([]byte, 0, len(data))
	for i := 0; i < len(data); i += 117 {
		if i+117 < len(data) {
			partial, err1 := rsa.EncryptPKCS1v15(rand.Reader, pub, data[i:i+117])
			if err1 != nil {
				return nil
			}
			encrypted = append(encrypted, partial...)
		} else {
			partial, err1 := rsa.EncryptPKCS1v15(rand.Reader, pub, data[i:])
			if err1 != nil {
				return nil
			}
			encrypted = append(encrypted, partial...)
		}
	}
	return encrypted
}

//DecryptRSA decrypt given []byte with RSA algorithm
func DecryptRSA(data []byte) []byte {
	if data == nil {
		return nil
	}
	block, _ := pem.Decode([]byte(PrivateKey))
	if block == nil {
		return nil
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	priv := privInterface.(*rsa.PrivateKey)
	decrypted := make([]byte, 0, len(data))
	for i := 0; i < len(data); i += 128 {
		if i+128 < len(data) {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, priv, data[i:i+128])
			if err1 != nil {
				return nil
			}
			decrypted = append(decrypted, partial...)
		} else {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, priv, data[i:])
			if err1 != nil {
				return nil
			}
			decrypted = append(decrypted, partial...)
		}
	}
	return decrypted
}

//SignWithRSA sign given encrypted data with RSA algorithm
func SignWithRSA(raw []byte) []byte {
	if raw == nil {
		return nil
	}
	data := EncryptMD5(EncryptSHA(raw))
	block, _ := pem.Decode([]byte(PrivateKey))
	if block == nil {
		return nil
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	priv := privInterface.(*rsa.PrivateKey)
	signed, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, data)
	if err != nil {
		return nil
	}
	return signed
}

//VerifySignature verify whether the given signature is correct
func VerifySignature(raw []byte, signature string) bool {
	if raw == nil || signature == "" {
		return false
	}
	data := EncryptMD5(EncryptSHA(raw))
	block, _ := pem.Decode([]byte(PublicKey))
	if block == nil {
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	pub := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.MD5, data, DecryptBase64(signature))
	if err != nil {
		return false
	}
	return true
}
