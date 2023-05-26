package main

import (
	"fmt"
	"godesignpattern/structural_pattern/decorator"
)

func main() {
	pizza := &decorator.VeggeMania{}

	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	pizzaWithCheeseAndTomoto := &decorator.TomatoTopping{Pizza: pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n\n", pizzaWithCheeseAndTomoto.GetPrice())
}
