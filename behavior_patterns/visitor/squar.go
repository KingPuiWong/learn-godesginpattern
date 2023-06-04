package visitor

type Square struct {
	Side int
}

func (s *Square) GetType() string {
	return "square"
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}
