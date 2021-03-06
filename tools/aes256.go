package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

func AesEncode(data []byte, key []byte) []byte {
	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	if err != nil {
		Alert("Error", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		Alert("Error", err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	ciphetText := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphetText

}
func AesDecode(encryptData []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		Alert("Error", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		Alert("Error", err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := encryptData[:nonceSize], encryptData[nonceSize:]
	data, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		Alert("Error", err.Error())
	}
	return data
}

func DataUnlock(data []byte, pwd string) ([]byte, error) {
	fmt.Println("dataUnlock")
	var key []byte
	pwdHash := sha256.New()
	i, err := pwdHash.Write([]byte(pwd))
	if err != nil {
		Alert("ERROR", err.Error())
		return data, err
	}
	fmt.Println(i)
	key = pwdHash.Sum(nil)
	result := AesDecode(data, key)
	return result, nil
}
func DataLock(data []byte, pwd string) ([]byte, error) {
	fmt.Println("dataLock")
	var key []byte
	pwdHash := sha256.New()
	i, err := pwdHash.Write([]byte(pwd))
	if err != nil {
		Alert("ERROR", err.Error())
	}
	fmt.Println(i)
	key = pwdHash.Sum(nil)
	return AesEncode(data, key), nil

}
