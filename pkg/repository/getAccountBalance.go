package repository

import (
	"context"
	"sqlBankCLI/pkg/errors"
)

func (r *Repository) GetAccountBalance(phoneNumber string) (balance float64, err error) {
	row := r.Conn.QueryRow(context.Background(),`
	SELECT 
		balance
	FROM 
		account
	WHERE 
		phone_number = $1
	`, phoneNumber)

	err = row.Scan(&balance)

	if err != nil {
		return 0, errors.ErrNoAccount
	}

	return

}