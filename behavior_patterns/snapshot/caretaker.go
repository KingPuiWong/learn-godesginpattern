package snapshot

//负责人

type Caretaker struct {
	MementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.MementoArray[index]
}
