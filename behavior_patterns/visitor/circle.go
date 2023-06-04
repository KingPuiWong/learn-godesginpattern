package visitor

type Circle struct {
	Radius int
}

func (c *Circle) GetType() string {
	return "circle"
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}
