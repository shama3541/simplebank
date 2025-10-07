package util

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(HashedPassword), nil
}

func CheckHashesPassword(password string, hasedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hasedpassword))
}
