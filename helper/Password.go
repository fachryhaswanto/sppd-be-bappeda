package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	var bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(passwordDb, passwordRequest string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(passwordDb), []byte(passwordRequest))
	if err != nil {
		return false
	} else {
		return true
	}

}
