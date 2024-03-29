package helper_auth_svc

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "problem at hashing ")
	}
	return string(HashedPassword)
}

func CompairPassword(hashedPassword string, plainPassword string) error {
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("passwords do not match")
	}

	return nil
}
