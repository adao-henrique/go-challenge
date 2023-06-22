package extensions

import "golang.org/x/crypto/bcrypt"

func HASH(secret string) (*string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(secret), 8)
	if err != nil {
		return nil, err
	}

	string_hashed := string(hashed)
	return &string_hashed, nil
}

func CompareHashAndValue(hash string, value string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
}
