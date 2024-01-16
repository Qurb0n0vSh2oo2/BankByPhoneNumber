package main

import (
	"fmt"
	"sqlBankCLI/pkg/config"
	"sqlBankCLI/pkg/database"
	"sqlBankCLI/pkg/repository"
	"sqlBankCLI/pkg/service"
	"sqlBankCLI/pkg/transport"
)

func main() {
	conf := config.NewConfig()
	db := database.NewDatabase(conf)
	repo := repository.NewRepository(db, *conf)
	svc := service.NewService(repo, *conf)
	transp := transport.NewTransport(svc)

	for {
		var choice int

		fmt.Println("<<BANK CLI WITH SQL>>")
		fmt.Println("1. Создать счет в банке")
		fmt.Println("2. Снять деньги")
		fmt.Println("3. Пополнение кошелька")
		fmt.Println("4. Перевод")
		fmt.Println("5. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			transp.CreateBankAccount()
		} else if choice == 2 {
			transp.WithdrawBankAccount()
		} else if choice == 3 {
			transp.TopUpBankAccount()
		} else if choice == 4 {
			transp.TransferMoneyByPhoneNumber()
		} else if choice == 5 {
			return
		}
	}
}
