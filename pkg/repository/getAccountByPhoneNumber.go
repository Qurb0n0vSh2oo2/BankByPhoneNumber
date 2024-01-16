package repository

import (
	"context"
	"sqlBankCLI/pkg/errors"
	"sqlBankCLI/pkg/models"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetAccountByPhoneNumber(PhoneNumber string) (account models.Account, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
		SELECT id, full_name, phone_number, address, balance FROM account WHERE phone_number = $1
	`, PhoneNumber)

	err = row.Scan(
		&account.Id,
		&account.FullName,
		&account.PhoneNumber,
		&account.Address,
		&account.Balance,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return account, errors.ErrDataNotFound
		}
	}

	return
}
