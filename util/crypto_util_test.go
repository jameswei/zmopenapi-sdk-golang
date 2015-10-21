package util

import (
	"fmt"
	"testing"
)

func TestEncryptBase64(t *testing.T) {
	raw := "linkedin china"
	encrypted := EncryptBase64([]byte(raw))
	fmt.Printf("'%s' encrypt using base64 is '%s'", raw, encrypted)
	if encrypted != "bGlua2VkaW4gY2hpbmE=" {
		t.Errorf("EncrypeBase64 failed, result is '%s', expected '%s'", encrypted, "bGlua2VkaW4gY2hpbmE=")
	}
}

func TestDecryptBase64(t *testing.T) {
	encrypted := "bGlua2VkaW4gY2hpbmE="
	decrypted := string(DecryptBase64(encrypted))
	fmt.Printf("'%s' decrypt using base64 is '%s'", encrypted, decrypted)
	if decrypted != "linkedin china" {
		t.Error("DecryptBase64 failed, result is '%s',expected '%s'", decrypted, "linkedin china")
	}
}

func TestEncryptMD5(t *testing.T) {
	if data := EncryptBase64(EncryptMD5([]byte("golang"))); data != "IcwoQJcpVl/BpNLdktsmnw==" {
		t.Error("EncryptMD5 failed")
	}
}

func TestEncryptSHA(t *testing.T) {
	if data := EncryptBase64(EncryptSHA([]byte("golang"))); data != "IcwoQJcpVl/BpNLdktsmnw==" {
		t.Error("EncryptSHA failed")
	}
}

func TestEncryptRSA(t *testing.T) {
	raw := "linkedin china"
	encryptedByPublicKey := EncryptBase64(EncryptRSA([]byte(raw)))
	fmt.Println(encryptedByPublicKey)
}
