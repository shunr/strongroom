package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Panic(err)
	}
	return key, &key.PublicKey
}
