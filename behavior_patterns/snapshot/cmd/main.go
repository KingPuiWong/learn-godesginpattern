package main

import (
	"fmt"
	"godesignpattern/behavior_patterns/snapshot"
)

func main() {
	caretaker := &snapshot.Caretaker{
		MementoArray: make([]*snapshot.Memento, 0),
	}

	originator := snapshot.Originator{
		State: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State:%s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State:%s\n", originator.GetState())

}
