package credential

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/gob"

	"github.com/deadblue/elevengo"
)

func decrypt(src []byte, secret string) []byte {
	block, _ := aes.NewCipher(deriveKey(secret))
	blockSize := block.BlockSize()

	plaintext := make([]byte, len(src)-blockSize)

	dec := cipher.NewCBCDecrypter(block, src[:blockSize])
	dec.CryptBlocks(plaintext, src[blockSize:])

	return plaintext
}

func Decode(src []byte, secret string, cred *elevengo.Credential) (err error) {
	// Prepare buffer
	decodeSize := base64.StdEncoding.DecodedLen(len(src))
	cookieData := make([]byte, decodeSize)
	// Base64 decode
	if decodeSize, err = base64.StdEncoding.Decode(cookieData, src); err != nil {
		return err
	} else {
		cookieData = cookieData[:decodeSize]
	}
	// Decrypt
	if secret != "" {
		cookieData = decrypt(cookieData, secret)
	}
	// GOB decode
	dec := gob.NewDecoder(bytes.NewReader(cookieData))
	return dec.Decode(cred)
}
