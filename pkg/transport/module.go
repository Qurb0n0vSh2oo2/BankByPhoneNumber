package transport

import "sqlBankCLI/pkg/service"

type Transport struct {
	Svc service.ServiceInterface
}

type TransportInterface interface {
	WithdrawBankAccount()
	TopUpBankAccount()
	CreateBankAccount()
	TransferMoneyByPhoneNumber()
}

func NewTransport(svc service.ServiceInterface) TransportInterface {
	return &Transport{
		Svc: svc,
	}
}
