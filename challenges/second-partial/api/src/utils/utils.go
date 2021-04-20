package utils

import (
  "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Remove(slice []string, val string) []string {
    for i, curr := range slice {
        if curr == val {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}
