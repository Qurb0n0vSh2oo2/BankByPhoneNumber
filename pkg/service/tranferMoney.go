package service

import (
	"fmt"
	"sqlBankCLI/pkg/errors"
	"sqlBankCLI/pkg/models"
)

func (s *Service) TransferMoney(sender models.Account, recipient models.Account, amount float64) (err error) {
	senderrow, err := s.Repo.GetAccountByPhoneNumber(sender.PhoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}
	fmt.Println(senderrow.Id)

	recipientrow, err := s.Repo.GetAccountByPhoneNumber(recipient.PhoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}
	
	fmt.Println(recipientrow.Id)
	if sender.PhoneNumber == recipient.PhoneNumber {
		return errors.ErrSamePhoneNumber
	}

	balance, err := s.Repo.GetAccountBalance(sender.PhoneNumber)
	if balance < amount {
		return errors.ErrNotEnoughMoney
	}

	var comission float64 = amount / 100 * s.Repo.GetPercent()
	fmt.Println("1")
	s.Repo.TransferMoneyFromSender(sender.PhoneNumber, amount)
	fmt.Println("2")
	s.Repo.TransferMoneyToRecipient(recipient.PhoneNumber, (amount - comission))
	fmt.Println("3")
	s.Repo.TopUpProfitAccount(comission)
	fmt.Println("4")
	s.Repo.Transfer(senderrow, recipientrow, amount)
	fmt.Println("5")
	return errors.ErrSuccess
}
