package login

import (
	"context"
	"errors"
	"testing"

	"github.com/adao-henrique/go-challenge/domain/entities"
	"github.com/adao-henrique/go-challenge/extensions"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogin(t *testing.T) {
	t.Run("Should create a valid token", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		secret := "secret_key"
		cpf := "818.499.480-03"

		acc, _ := entities.NewAccount("name test", cpf, secret)
		customClaims := map[string]string{"account_origin_id": acc.ID}

		inputInit := InputInit{
			GetByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return acc, nil
			},
		}

		loginService := Init(inputInit)
		token, err := loginService.Login(ctx, LoginUserInput{Cpf: cpf, Secret: secret})
		assert.Nil(t, err)

		result, err := extensions.ValidateToken(token, customClaims)
		assert.Nil(t, err)
		assert.True(t, result)
	})

	t.Run("CPF Error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		secret := "secret_key"
		cpf := "121.111.111-12"

		acc, _ := entities.NewAccount("name test", cpf, secret)

		inputInit := InputInit{
			GetByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return acc, nil
			},
		}

		loginService := Init(inputInit)
		_, err := loginService.Login(ctx, LoginUserInput{Cpf: cpf, Secret: secret})
		assert.Equal(t, err, errors.New("invalid CPF"))
	})

	t.Run("Inorrect secret", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		secret := "secret_key"
		incorrectSecret := "incorrectSecret"
		cpf := "818.499.480-03"

		acc, _ := entities.NewAccount("name test", cpf, secret)

		inputInit := InputInit{
			GetByCPFMock: func(ctx context.Context, cpf string) (entities.Account, error) {
				return acc, nil
			},
		}

		loginService := Init(inputInit)
		_, err := loginService.Login(ctx, LoginUserInput{Cpf: cpf, Secret: incorrectSecret})

		assert.Equal(t, err, errors.New("erro comparar secrets"))
	})
}
