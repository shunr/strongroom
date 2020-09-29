package util

func Bxor(b1 []byte, b2 []byte) []byte {
	result := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		result[i] = b1[i] ^ b2[i]
	}
	return result
}
