package main

import (
	"fmt"
	"godesignpattern/structural_pattern/face"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := face.NewWalletFacade("abc", 1235)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
