package crypto

import (
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/scrypt"
)

func HKDF(secret []byte, salt []byte, info []byte, key_length int) []byte {
	hash_fn := sha256.New
	hash := hkdf.New(hash_fn, secret, salt, info)
	key := make([]byte, key_length)
	if _, err := io.ReadFull(hash, key); err != nil {
		panic(err)
	}
	return key
}

func Scrypt(secret []byte, salt []byte, key_length int) []byte {
	// Using N = 32768, see https://cryptobook.nakov.com/mac-and-key-derivation/scrypt
	dk, err := scrypt.Key(secret, salt, 1<<15, 8, 1, key_length)
	if err != nil {
		panic(err)
	}
	return dk
}
