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
	"linkedin/ginutil"
)

const (
	PrivateKey = `-----BEGIN PRIVATE KEY-----
PUT YOUR PRIVATE KEY
-----END PRIVATE KEY-----`
	PublicKey = `-----BEGIN PUBLIC KEY-----
PUT ZHIMA PUBLIC KEY
-----END PUBLIC KEY-----`
	TestPrivateKey = `-----BEGIN PRIVATE KEY-----
PUT YOUR TEST PRIVATE KEY
-----END PRIVATE KEY-----`
	TestPublicKey = `-----BEGIN PUBLIC KEY-----
PUT ZHIMA TEST PUBLIC KEY
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
	publicKey := []byte(PublicKey)
	if !ginutil.IsProduction() {
		publicKey = []byte(TestPublicKey)
	}
	block, _ := pem.Decode(publicKey)
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
	privateKey := []byte(PrivateKey)
	if !ginutil.IsProduction() {
		privateKey = []byte(TestPrivateKey)
	}
	block, _ := pem.Decode(privateKey)
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
func SignRSA(raw []byte, algorithm crypto.Hash) []byte {
	if raw == nil {
		return nil
	}
	privateKey := []byte(PrivateKey)
	if !ginutil.IsProduction() {
		privateKey = []byte(TestPrivateKey)
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	priv := privInterface.(*rsa.PrivateKey)
	var data []byte
	if algorithm == crypto.SHA1 {
		data = EncryptSHA(raw)
	} else {
		data = EncryptMD5(EncryptSHA(raw))
	}
	signed, err := rsa.SignPKCS1v15(rand.Reader, priv, algorithm, data)
	if err != nil {
		return nil
	}
	return signed
}

//VerifySignature verify whether the given signature is correct
func VerifySignature(raw []byte, signature string, algorithm crypto.Hash) bool {
	if raw == nil || signature == "" {
		return false
	}
	publicKey := []byte(PublicKey)
	if !ginutil.IsProduction() {
		publicKey = []byte(TestPublicKey)
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	pub := pubInterface.(*rsa.PublicKey)
	var data []byte
	if algorithm == crypto.SHA1 {
		data = EncryptSHA(raw)
	} else {
		data = EncryptMD5(EncryptSHA(raw))
	}
	err = rsa.VerifyPKCS1v15(pub, algorithm, data, DecryptBase64(signature))
	if err != nil {
		return false
	}
	return true
}
