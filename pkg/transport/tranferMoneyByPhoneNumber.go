package transport

import (
	"fmt"
)

func (t *Transport) TransferMoneyByPhoneNumber() {
	var senderPhoneNumber, recipientPhoneNumber string
	var amount float64

	fmt.Println("Введите номер отправителя: ")
	fmt.Scan(&senderPhoneNumber)

	fmt.Println("Введите номер получателя: ")
	fmt.Scan(&recipientPhoneNumber)

	fmt.Println("Введите сумму для первода: ")
	fmt.Scan(&amount)

	err := t.Svc.TransferMoney(senderPhoneNumber, recipientPhoneNumber, amount)
	if err != nil {
		fmt.Println(err)
		return
	}

}
