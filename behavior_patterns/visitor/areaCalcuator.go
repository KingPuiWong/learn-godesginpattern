package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a AreaCalculator) VisitForSquare(square *Square) {
	fmt.Println("calculating area for square")
}

func (a AreaCalculator) VisitForCircle(circle *Circle) {
	fmt.Println("calculating area for circle")
}
