package utils

import "golang.org/x/crypto/bcrypt"

func CreateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

func CheckHash(str, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str))
	return err == nil
}
