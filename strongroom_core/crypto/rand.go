package crypto

import (
	"crypto/rand"
)

func RandBytes(length int) []byte {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}
