package face

import "fmt"

type Wallet struct {
	Balance int
}

func newWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.Balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.Balance < amount {
		return fmt.Errorf("blance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.Balance = w.Balance - amount
	return nil
}
