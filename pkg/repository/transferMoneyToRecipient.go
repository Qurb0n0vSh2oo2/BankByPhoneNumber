package repository

import "context"

func (r *Repository) TransferMoneyToRecipient(senderPhoneNumber string, amount float64) (err error) {
	_, err = r.Conn.Exec(context.Background(), `
	UPDATE account
	SET balance = balance + $1
	WHERE phone_number = $2
	`, amount, senderPhoneNumber)

	if err != nil {
		return err
	}
	return
}
