package main

import "godesignpattern/behavior_patterns/mediator"

//中介者模式: 也成为控制器模式，能让你减少对象之间混乱无序的依赖关系。该模式会限制对象之间的直接交互‘
//迫使它们通过一个中介者对象进行合作

func main() {
	stationManager := mediator.NewStationManger()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}

	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()

}
