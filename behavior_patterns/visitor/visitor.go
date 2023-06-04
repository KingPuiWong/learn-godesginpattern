package visitor

type Visitor interface {
	VisitForSquare(square *Square)
	VisitForCircle(circle *Circle)
}
