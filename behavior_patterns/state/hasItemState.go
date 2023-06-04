package state

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h HasItemState) AddItem(i int) error {
	fmt.Printf("%d items added\n", i)
	h.vendingMachine.incrementItemCount(i)
	return nil
}

func (h HasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("item requested\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (h HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
