package factory

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{Gun{
		name:  "Ak47 gun",
		power: 4,
	}}
}
