package errors

import "errors"

var (
	ErrIncorrectPhoneNumber = errors.New("длина номера телефона не равняется 9")
	ErrDataNotFound         = errors.New("не найдено")
	ErrAlreadyHasAccount    = errors.New("у вас уже есть счет в нашем банке")
	ErrNoAccount            = errors.New("извините, но у вас нету счёта")
	ErrNotEnoughMoney       = errors.New("недостаточно средств")
	ErrSuccess              = errors.New("Успешно переведено")
	ErrSamePhoneNumber      = errors.New("вы не можете себе отправить деньги, для этого выберите раздел пополнение кошелька")
)
