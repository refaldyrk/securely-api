package helper

import (
	"github.com/mervick/aes-everywhere/go/aes256"
	"github.com/spf13/viper"
)

func EncryptAES(s string) string {
	return aes256.Encrypt(s, viper.GetString("AES_SECRET_KEY"))
}

func DecryptAES(s string) string {
	return aes256.Decrypt(s, viper.GetString("AES_SECRET_KEY"))
}

func EncryptAES2(s string, p string) string {
	return aes256.Encrypt(s, p)
}

func DecryptAES2(s string, p string) string {
	return aes256.Decrypt(s, p)
}
