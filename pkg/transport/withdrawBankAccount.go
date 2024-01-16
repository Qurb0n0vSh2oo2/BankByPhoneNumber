package transport

import (
	"fmt"
	"sqlBankCLI/pkg/models"
)

func (t *Transport) WithdrawBankAccount() {
	var account models.Account
	var amount float64

	fmt.Println("Введите номер телефона:")
	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Какую сумму хотите снять:")
	fmt.Scan(&amount)

	err := t.Svc.WithdrawBankAccount(account.PhoneNumber, amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Уcпешно снято!")
}
