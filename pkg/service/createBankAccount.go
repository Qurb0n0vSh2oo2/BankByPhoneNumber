package service

import (
	"sqlBankCLI/pkg/errors"
	"sqlBankCLI/pkg/models"
)

func (s *Service) CreateBankAccount(account models.Account) error {
	if len(account.PhoneNumber) != 9 {
		return errors.ErrIncorrectPhoneNumber
	}

	account2, err := s.Repo.GetAccountByPhoneNumber(account.PhoneNumber)

	// if err != errors.ErrDataNotFound {
	// 	return err
	// }

	if account2.PhoneNumber == account.PhoneNumber {
		return errors.ErrAlreadyHasAccount
	}

	err = s.Repo.CreateBankAccount(account)

	if err != nil {
		return err
	}

	return nil
}
