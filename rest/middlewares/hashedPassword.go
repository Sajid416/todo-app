package middlewares

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(pass), 
		bcrypt.DefaultCost,
	)
	return string(hashedPass), err

}
