package service

import (
	"sqlBankCLI/pkg/config"
	"sqlBankCLI/pkg/models"
	"sqlBankCLI/pkg/repository"
)

type Service struct {
	Repo    repository.RepositoryInterface
	Percent config.Config
}

type ServiceInterface interface {
	CreateBankAccount(account models.Account) error
	WithdrawBankAccount(phoneNumber string, amount float64) (err error)
	TopUpBankAccount(phoneNumber string, amount float64) (err error)
	TransferMoney(sender models.Account, recipient models.Account, amount float64) (err error)
}

func NewService(repo repository.RepositoryInterface, percent config.Config) ServiceInterface {
	return &Service{
		Repo: repo,
		Percent: percent,
	}
}
