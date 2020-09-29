package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func AESGCMEncrypt(data []byte, key []byte, nonce []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return aesgcm.Seal(nil, nonce, data, nil)
}

func AESGCMDecrypt(ciphertext []byte, key []byte, nonce []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	decrypted, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return decrypted
}
