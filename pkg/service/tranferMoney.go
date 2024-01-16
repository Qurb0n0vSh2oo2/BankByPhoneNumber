package service

import (
	"sqlBankCLI/pkg/errors"
)

func (s *Service) TransferMoney(senderPhoneNumber, recipientPhoneNumber string, amount float64) (err error) {
	_, err = s.Repo.GetAccountByPhoneNumber(senderPhoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}

	_, err = s.Repo.GetAccountByPhoneNumber(recipientPhoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}

	balance, err := s.Repo.GetAccountBalance(senderPhoneNumber)
	if balance < amount {
		return errors.ErrNotEnoughMoney
	}
	var comission float64 = amount / 100 * s.Repo.GetPercent()
	s.Repo.TransferMoneyFromSender(senderPhoneNumber, amount)
	s.Repo.TransferMoneyToRecipient(recipientPhoneNumber, (amount - comission))
	s.Repo.TopUpProfitAccount(comission)
	s.Repo.Transfer(senderPhoneNumber, recipientPhoneNumber, amount)
	return errors.ErrSuccess
}
