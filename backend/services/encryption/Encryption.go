package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var Key = "AIRLIFTUSAHASHKEYENCRYPT%*@!4567"

func init() {
}

// Encrypt func used to Encrypt data
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return base64.StdEncoding.EncodeToString([]byte(ciphertext)), nil
}

// Decrypt func used to Decrypt text
func Decrypt(text string) (string, error) {
	encodeString, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	value := encodeString
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	if len(value) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := value[:aes.BlockSize]
	value = value[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(value, value)
	data, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
