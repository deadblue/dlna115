package credential

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"

	"github.com/deadblue/elevengo"
)

func encrypt(src []byte, secret string) []byte {
	block, _ := aes.NewCipher(deriveKey(secret))
	blockSize := block.BlockSize()

	// Padding
	padSize := blockSize - len(src)%blockSize
	padding := make([]byte, padSize)
	for i := 0; i < padSize; i++ {
		padding[i] = byte(padSize)
	}
	plaintext := append(src, padding...)

	// Prepare ciphertext buffer
	ciphertext := make([]byte, blockSize+len(plaintext))

	// Generate IV
	rand.Read(ciphertext[:blockSize])

	// Encrypt
	enc := cipher.NewCBCEncrypter(block, ciphertext[:blockSize])
	enc.CryptBlocks(ciphertext[blockSize:], plaintext)
	return ciphertext
}

func Encode(cred *elevengo.Credential, secret string) string {
	// Marshal credential by gob
	buf := &bytes.Buffer{}
	gobEnc := gob.NewEncoder(buf)
	gobEnc.Encode(cred)

	cookieData := buf.Bytes()
	if secret != "" {
		cookieData = encrypt(cookieData, secret)
	}

	return base64.StdEncoding.EncodeToString(cookieData)
}
