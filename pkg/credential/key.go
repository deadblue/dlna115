package credential

import "crypto/sha256"

func deriveKey(secret string) []byte {
	h := sha256.New()
	h.Write([]byte(secret))
	return h.Sum(nil)
}
