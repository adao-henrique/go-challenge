package login

import (
	"errors"

	"github.com/adao-henrique/go-challenge/extensions"
)

func (ls LoginService) Autenticate(token string) (interface{}, error) {
	claims, err := extensions.GetClaims(token)
	if err != nil {
		return "", err
	}

	key := "account_origin_id"
	if id := claims[key]; id != nil {
		return id, nil
	}

	return nil, errors.New("invalid token")
}
