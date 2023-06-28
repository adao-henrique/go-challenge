package login

import (
	"context"
	"errors"
	"fmt"

	"github.com/adao-henrique/go-challenge/extensions"
)

func (ls LoginService) Login(ctx context.Context, input LoginUserInput) (string, error) {

	validCPF := extensions.ValidateCPF(input.Cpf)
	if !validCPF {
		fmt.Println("validCPF: ", validCPF)

		return "", errors.New("invalid CPF")
	}

	account, err := ls.accountUseCases.GetByCPF(ctx, input.Cpf)
	if err != nil {
		return "", err
	}

	err = extensions.CompareHashAndValue(account.Secret, input.Secret)
	if err != nil {
		return "", errors.New("erro comparar secrets")
	}

	customClaims := map[string]interface{}{
		"account_origin_id": account.ID,
	}

	return extensions.NewJWT(customClaims)
}
