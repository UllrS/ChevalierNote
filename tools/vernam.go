package tools

import (
	"math/rand"
	"time"
)

func VernamEncode(data []byte) ([]byte, []byte) {
	key_len := len(data)
	var key []byte
	var encryptedData []byte

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < key_len; i++ {
		random := rand.Intn(256)
		key = append(key, byte(random))
		encryptedData = append(encryptedData, byte((int(key[i])+int(data[i]))%255))
	}
	return encryptedData, key

}
func VernamDencode(key []byte, encryptedData []byte) []byte {
	key_len := len(encryptedData)
	var data []byte
	for i := 0; i < key_len; i++ {
		b := (int(data[i]) - int(key[i])) % 255
		if b < 0 {
			b = (b + 255) % 255
		} else {
			b = b % 255
		}
		data = append(data, byte(b))
	}
	return data
}
