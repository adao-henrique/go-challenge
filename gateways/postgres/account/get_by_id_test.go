package account

// func TestCreateAccouunt(t *testing.T) {
// 	t.Run("Should create an account", func(t *testing.T) {
// 		t.Parallel()
// 		ctx := context.Background()

// 		expectAccount, err := entities.NewAccount("name test", "392.580.620-27", "secret_test")
// 		assert.Nil(t, err)

// 		sql := `insert into public.account(id, name, cpf, balance, created_at) values ($1,$2,$3,$4,$5)`
// 		_, err = db.Exec(
// 			ctx,
// 			sql,
// 			expectAccount.ID,
// 			expectAccount.Name,
// 			expectAccount.Cpf,
// 			expectAccount.Balance,
// 			expectAccount.CreatedAt,
// 		)
// 		assert.Nil(t, err)

// 		repository, err := NewRepository(*db)
// 		assert.Nil(t, err)

// 		account, err := repository.GetByID(ctx, expectAccount.ID)
// 		assert.Nil(t, err)

// 		assert.EqualValues(t, expectAccount, account)
// 	})
// }
