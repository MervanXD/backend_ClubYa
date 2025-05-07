package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// Genera una clave de 32 bytes desde un string (clave maestra)
func generateKey(secret string) []byte {
	hash := sha256.Sum256([]byte(secret))
	return hash[:]
}

// Desencripta una cadena base64 cifrada con AES-256 CBC
func DecryptAES(encryptedBase64, secret string) (string, error) {
	key := generateKey(secret)

	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	if len(encryptedData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := encryptedData[:aes.BlockSize]
	cipherText := encryptedData[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	// Quitar padding
	plainText, err := pkcs7Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// Desencriptador de entorno
func DecryptEnv(encrypted string) (string, error) {
	const secretKey = "[Code Here]" // tu clave secreta
	return DecryptAES(encrypted, secretKey)
}

// ----------------------
// Padding Helpers (PKCS7)
// ----------------------

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 || length%blockSize != 0 {
		return nil, errors.New("invalid padding size")
	}

	padLen := int(data[length-1])
	if padLen == 0 || padLen > blockSize {
		return nil, errors.New("invalid padding")
	}

	for _, v := range data[length-padLen:] {
		if int(v) != padLen {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:length-padLen], nil
}
