package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Panic(err)
	}
	return key, &key.PublicKey
}

func EncryptRSAOAEP(plaintext []byte, label []byte, public_key *rsa.PublicKey) ([]byte, error) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, public_key, plaintext, label)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func DecryptRSAOAEP(ciphertext []byte, label []byte, private_key *rsa.PrivateKey) ([]byte, error) {
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, private_key, ciphertext, label)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
