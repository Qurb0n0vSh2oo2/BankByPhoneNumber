package repository

import (
	"context"
)

func (r *Repository) Transfer(sender, recipient string, amount float64) (err error) {
	_, err = r.Conn.Exec(context.Background(), `
	INSERT 
	INTO 
		tranfer(sender_id, recipient_id, amount)
	VALUES
		($1, $2, $3)`,
		sender, recipient, amount)

	if err != nil {
		return err
	}
	return
}
