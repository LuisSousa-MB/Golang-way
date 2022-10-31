package main

import "fmt"

var OvenTime int = 40

func main() {
	remain := RemainingOvenTime(10)
	fmt.Printf("Falta %d minutos\n", remain)
	timePreparation := PreparationTime(8)
	fmt.Printf("Tempo de preparação foi de %d minutos\n", timePreparation)
	elapsedTime := ElapsedTime(8, 10)
	fmt.Printf("Tempo passado até agora é de %d minutos\n", elapsedTime)

}

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}
