package transport

import (
	"fmt"
	"sqlBankCLI/pkg/models"
)

func (t *Transport) TransferMoneyByPhoneNumber() {
	var sender, recipient models.Account
	var amount float64

	fmt.Println("Введите номер отправителя: ")
	fmt.Scan(&sender.PhoneNumber)

	fmt.Println("Введите номер получателя: ")
	fmt.Scan(&recipient.PhoneNumber)

	fmt.Println("Введите сумму для первода: ")
	fmt.Scan(&amount)

	err := t.Svc.TransferMoney(sender, recipient, amount)
	if err != nil {
		fmt.Println(err)
		return
	}

}
