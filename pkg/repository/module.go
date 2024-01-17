package repository

import (
	"sqlBankCLI/pkg/config"
	"sqlBankCLI/pkg/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn    *pgx.Conn
	Percent config.Config
}

type RepositoryInterface interface {
	CreateBankAccount(account models.Account) (err error)
	GetAccountByPhoneNumber(PhoneNumber string) (account models.Account, err error)
	GetAccountBalance(phoneNumber string) (balance float64, err error)
	WithdrawBankAccount(phoneNumber string, amount float64) (err error)
	TopUpBankAccount(phoneNumber string, amount float64) (err error)
	TransferMoneyFromSender(senderPhoneNumber string, amount float64) (err error)
	TransferMoneyToRecipient(senderPhoneNumber string, amount float64) (err error)
	GetPercent() float64
	TopUpProfitAccount(amount float64)
	Transfer(sender models.Account, recipient models.Account, amount float64) (err error)
}

func NewRepository(conn *pgx.Conn, percent config.Config) *Repository {
	return &Repository{
		Conn:    conn,
		Percent: percent,
	}
}
