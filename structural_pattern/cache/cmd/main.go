package main

import (
	"fmt"
	"godesignpattern/structural_pattern/cache"
)

func main() {
	game := cache.NewGame()

	game.AddTerrorist(cache.TerroristDressType)
	game.AddTerrorist(cache.TerroristDressType)
	game.AddTerrorist(cache.TerroristDressType)
	game.AddTerrorist(cache.TerroristDressType)

	game.AddCounterTerrorist(cache.CounterTerrorristDressType)
	game.AddCounterTerrorist(cache.CounterTerrorristDressType)
	game.AddCounterTerrorist(cache.CounterTerrorristDressType)
	game.AddCounterTerrorist(cache.CounterTerrorristDressType)

	dressFactoryInstance := cache.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
