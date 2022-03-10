package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

func main() {
	b := []byte{00, 215, 21}
	d := PwdTest(b, "1234")
	fmt.Println("d")
	fmt.Println(d)
	c := PwdTestUnlock(d, "123")
	fmt.Println("c")
	fmt.Println(c)

}

func PwdTest(data []byte, pwd string) []byte {
	fmt.Println("LOCK")

	var unlockData []byte
	unlockData = DataLock(data, pwd)
	fmt.Println("LOCK END")
	return unlockData
}
func PwdTestUnlock(data []byte, pwd string) []byte {
	var unlockData []byte
	fmt.Println("UNLOCK")
	unlockData = DataUnlock(data, pwd)
	return unlockData

}
func DataUnlock(data []byte, pwd string) []byte {
	fmt.Println("dataUnlock")
	var key []byte
	pwdHash := sha256.New()
	_, err := pwdHash.Write([]byte(pwd))
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	key = pwdHash.Sum(nil)
	return AesDecode(data, key)
}
func DataLock(data []byte, pwd string) []byte {
	fmt.Println("dataLock")
	var key []byte
	pwdHash := sha256.New()
	_, err := pwdHash.Write([]byte(pwd))
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	key = pwdHash.Sum(nil)
	return AesEncode(data, key)

}

func AesEncode(data []byte, key []byte) []byte {
	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	ciphetText := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphetText

}
func AesDecode(encryptData []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := encryptData[:nonceSize], encryptData[nonceSize:]
	data, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("ERROR Bad pwd", err.Error())
	}
	return data
}
