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
    PrivateKey   = ``
	PublicKey    = ``
)

//EncryptBase64 encrypt given []byte with Base64 algorithm
func EncryptBase64(data []byte) string {
	encrypted := base64.StdEncoding.EncodeToString(data)
	return encrypted
}

//DecryptBase64 decrypt given string with Base64 algorithm
func DecryptBase64(data string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(data)
	return decrypted
}

//EncryptMD5 encrypt given []byte with MD5 algorithm
func EncryptMD5(data []byte) []byte {
	encrypter := md5.New()
	encrypter.Write(data)
	return encrypter.Sum(nil)
}

//EncryptSHA encrypt given []byte with SHA algorithm
func EncryptSHA(data []byte) []byte {
	encypter := sha1.New()
	encypter.Write(data)
	return encypter.Sum(nil)
}

//EncryptRSA encrypt given data with RSA algorithm
func EncryptRSA(data []byte) []byte {
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
			partial, err1 := rsa.EncryptPKCS1v15(rand.Reader, pub, data[i:len(data)])
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
	block, _ := pem.Decode([]byte(PrivateKey))
	if block == nil {
		return nil
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	decrypted := make([]byte, 0, len(data))
	for i := 0; i < len(data); i += 128 {
		if i+128 < len(data) {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), data[i:i+128])
			if err1 != nil {
				return nil
			}
			decrypted = append(decrypted, partial...)
		} else {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), data[i:len(data)])
			if err1 != nil {
				return nil
			}
			decrypted = append(decrypted, partial...)
		}
	}
	return decrypted
}

//SignWithRSA sign given encrypted data with RSA algorithm
func SignWithRSA(data []byte) []byte {
	block, _ := pem.Decode([]byte(PrivateKey))
	if block == nil {
		return nil
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	signed, err := rsa.SignPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), crypto.MD5, EncryptMD5(data))
	if err != nil {
		return nil
	}
	return signed
}

func VerifySignature(data []byte, signature string) bool {
	block, _ := pem.Decode([]byte(PublicKey))
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	pub := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pub, crypto.MD5, EncryptMD5(data), []byte(signature))
	if err != nil {
		return false
	}
	return true
}
