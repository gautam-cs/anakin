package config

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"github.com/golang-jwt/jwt"
)

//--- config ---
func accountJWTKey() string {
	if IsDebugEnv() {
		return appConfig.Account.JwtSecret
	}

	if appConfig.Account != nil {
		return appConfig.Account.JwtSecret
	}

	appConfig.Account = new(accountConfig)

	readDecodedSecret("account", appConfig.Account)

	return appConfig.Account.JwtSecret
}

func accountCryptKey() string {
	if IsDebugEnv() {
		return appConfig.Account.CryptKey
	}

	if appConfig.Account != nil {
		return appConfig.Account.CryptKey
	}

	appConfig.Account = new(accountConfig)

	readDecodedSecret("account", appConfig.Account)

	return appConfig.Account.CryptKey
}

//---exports
func DecodeSecretValue(encoded string) (string, error) {
	return decrypt(accountCryptKey(), encoded)
}

func EncodeSecretValue(plain string) (string, error) {
	return encrypt(accountCryptKey(), plain)
}

func EncodeJWTToken(claims map[string]interface{}) (string, error) {
	mClaims := jwt.MapClaims(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mClaims)
	return token.SignedString([]byte(accountJWTKey()))
}

// internal
func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encrypt(key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	msg := Pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	return finalMsg, nil
}

func decrypt(key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("block size must be multiple of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := Unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}
