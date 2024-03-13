package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	passByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passByte), err
}

func ComparePassword(password string, hasedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
	return err == nil
}
