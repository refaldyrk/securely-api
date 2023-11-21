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
