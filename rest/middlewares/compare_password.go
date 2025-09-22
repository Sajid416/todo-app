package middlewares

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Compare_Pass(userPass, hashPass string) bool {
	userPass = strings.TrimSpace(userPass)
	hashPass = strings.TrimSpace(hashPass)
	if userPass == "" || hashPass == "" {
		return false
	}
	fmt.Println(userPass)
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(userPass))
	return err == nil
}
