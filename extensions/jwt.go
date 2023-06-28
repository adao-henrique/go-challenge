package extensions

import (
	"errors"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func GetJWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)
}

func NewJWT(customClaims map[string]interface{}) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	tokenAuth := jwtauth.New("HS256", []byte(secret), nil)
	expiretIn := time.Now().Add(10 * time.Minute)

	customClaims["ext"] = expiretIn.Unix()
	_, tokenString, _ := tokenAuth.Encode(customClaims)

	return tokenString, nil
}

func ValidateToken(token string, customClaims map[string]string) (bool, error) {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)
	jwtToken, err := jwtauth.VerifyToken(tokenAuth, token)
	if err != nil {
		return false, err
	}

	claims := jwtToken.PrivateClaims()

	for key, v := range customClaims {
		if claims[key] != v {
			return false, errors.New("error to get jwt claims")
		}
	}

	return true, nil

}

func GetClaims(token string) (map[string]interface{}, error) {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)
	jwtToken, err := jwtauth.VerifyToken(tokenAuth, token)
	if err != nil {
		return nil, err
	}

	claims := jwtToken.PrivateClaims()
	return claims, nil
}
