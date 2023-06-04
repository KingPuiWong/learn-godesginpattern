package main

import (
	"fmt"
	"godesignpattern/behavior_patterns/state"
	"log"
)

func main() {
	vendingMachine := state.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem(2)

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}
