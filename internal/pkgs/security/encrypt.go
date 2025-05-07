package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encripta texto plano con AES-256 CBC y devuelve base64
func EncryptAES(plainText, secret string) (string, error) {
	key := generateKey(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedText := pkcs7Pad([]byte(plainText), aes.BlockSize)
	cipherText := make([]byte, aes.BlockSize+len(paddedText))
	iv := cipherText[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], paddedText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// ------------------------
// Padding Helpers (PKCS7)
// ------------------------

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytesRepeat(byte(padding), padding)
	return append(data, padText...)
}

func bytesRepeat(b byte, count int) []byte {
	buf := make([]byte, count)
	for i := range buf {
		buf[i] = b
	}
	return buf
}
