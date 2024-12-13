package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//the number 14 is how complex the hash should been
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	//he give us bytes, so we convert them to string
	return string(bytes), err
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	//the bcrypt only use byte, not string, so we have always to convert it
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
