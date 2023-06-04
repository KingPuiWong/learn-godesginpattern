package main

import (
	"fmt"
	"godesignpattern/behavior_patterns/visitor"
)

func main() {
	square := &visitor.Square{Side: 2}
	circle := &visitor.Circle{Radius: 3}

	areaCalculator := &visitor.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)

	fmt.Println()

	middleCoordinates := &visitor.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
}
