package middlewares

import "golang.org/x/crypto/bcrypt"

func Compare_Pass(userPass, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(userPass))
	return err == nil
}
