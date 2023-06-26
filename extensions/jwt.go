package extensions

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewJWT(customClaims map[string]string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(10 * time.Minute)

	for key, v := range customClaims {
		claims[key] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string, customClaims map[string]string) (bool, error) {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	// if err != nil {
	// 	fmt.Println(tkn, err)
	// 	return false, err
	// }

	for key, v := range customClaims {
		if claims[key] != v {
			return false, errors.New("Error to get jwt claims")
		}
	}

	return true, nil

}
