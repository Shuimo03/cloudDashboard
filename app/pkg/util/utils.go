package util

import (
	"github.com/containerd/containerd/log"
	"golang.org/x/crypto/bcrypt"
)

func BcryptPassWord(password string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.L.Error(err)
	}
	encodePassWord := string(pass)
	return encodePassWord
}
