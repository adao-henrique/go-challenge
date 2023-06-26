package login

import (
	"context"
	"errors"
	"fmt"

	"github.com/adao-henrique/go-challenge/extensions"
)

func (ls LoginService) login(ctx context.Context, input LoginUserInput) (string, error) {

	validCPF := extensions.ValidateCPF(input.cpf)
	if !validCPF {
		fmt.Println("validCPF: ", validCPF)

		return "", errors.New("invalid CPF")
	}

	account, err := ls.accountUseCases.GetByCPF(ctx, input.cpf)
	if err != nil {
		return "", err
	}

	err = extensions.CompareHashAndValue(account.Secret, input.secret)
	if err != nil {
		return "", errors.New("erro comparar secrets")
	}

	customClaims := map[string]string{
		"account_origin_id": account.ID,
	}

	return extensions.NewJWT(customClaims)
}
