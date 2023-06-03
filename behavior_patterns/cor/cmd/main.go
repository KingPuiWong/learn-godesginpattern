package main

import "godesignpattern/behavior_patterns/cor"

func main() {
	cashier := &cor.Cashier{}

	medical := &cor.Medical{}
	medical.SetNext(cashier)

	doctor := &cor.Doctor{}
	doctor.SetNext(medical)

	reception := &cor.Reception{}
	reception.SetNext(doctor)

	patient := &cor.Patient{Name: "abc"}
	reception.Execute(patient)
}
