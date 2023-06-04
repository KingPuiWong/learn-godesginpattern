package state

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i ItemRequestedState) AddItem(i2 int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i ItemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
