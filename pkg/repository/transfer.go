package repository

import (
	"context"
	"fmt"
	"sqlBankCLI/pkg/models"
)

func (r *Repository) Transfer(sender models.Account, recipient models.Account, amount float64) (err error) {
	fmt.Println(sender.Id)
	fmt.Println(recipient.Id)
	_, err = r.Conn.Exec(context.Background(), `
	INSERT INTO 
		transfer(sender_id, recipient_id, amount)
	VALUES
		($1, $2, $3)`,
		sender.Id, recipient.Id, amount)
	return
}

