package service

import "sqlBankCLI/pkg/errors"

func (s *Service) TopUpBankAccount(phoneNumber string, amount float64) (err error) {
	_, err = s.Repo.GetAccountByPhoneNumber(phoneNumber)
	if err != nil {
		return errors.ErrNoAccount
	}

	// balance, err := s.Repo.GetAccountBalance(phoneNumber)
	// if balance < amount {
	// 	return errors.ErrNotEnoughMoney
	// }

	s.Repo.TopUpBankAccount(phoneNumber, amount)
	return
}
