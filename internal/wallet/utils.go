package wallet

import (
	internalErr "github.com/anti-duhring/go-grpc/internal/errors"
	"github.com/anti-duhring/go-grpc/internal/schema"
)

func transferMoney(from *schema.Wallet, to *schema.Wallet, amount float64) error {

	if from.Balance < amount {
		return internalErr.New(internalErr.ErrWalletDontHaveEnoughMoney, "not enough money")
	}

	from.Balance -= amount
	to.Balance += amount

	return nil

}
