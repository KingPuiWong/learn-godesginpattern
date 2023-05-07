package main

import (
	"fmt"
	"godesignpattern/single"
)

func main() {
	for i := 0; i < 30; i++ {
		//go single.GetInstance()
		go single.GetInstanceAnotherExample()
	}
	fmt.Scanln()
}
