package repository

import (
	"context"
)

func (repo *Repository) WithdrawBankAccount(phoneNumber string, amount float64) (err error) {

	_, err = repo.Conn.Exec(context.Background(), `
	UPDATE account
	SET balance = balance - $1
	WHERE phone_number = $2
	`, amount, phoneNumber)

	return err
}
