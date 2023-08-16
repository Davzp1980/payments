package main

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Checlpassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
