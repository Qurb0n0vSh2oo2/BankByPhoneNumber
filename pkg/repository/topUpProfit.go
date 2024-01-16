package repository

import (
	"fmt"
	"sqlBankCLI/pkg/models"
)

func (repo *Repository) TopUpProfitAccount(amount float64) {
	var user models.Account
	user, err := repo.GetAccountByPhoneNumber("544")

	if err != nil {
		fmt.Println(err)
		user.FullName = "BankProfit"
		user.Balance = 0.0
		user.PhoneNumber = "544"
		err = repo.CreateBankAccount(user)
		if err != nil {
			return
		}
		user, _ = repo.GetAccountByPhoneNumber("544")
	}

	err = repo.TransferMoneyToRecipient(user.PhoneNumber, user.Balance+amount)
	if err != nil {
		return
	}
}
