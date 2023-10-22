package main

import "fmt"

// TODO: define the 'OvenTime' constant
const (
	OvenTime = 40
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var remainingOvenTime int
	remainingOvenTime = OvenTime - actualMinutesInOven
	//panic("RemainingOvenTime not implemented")
	return remainingOvenTime

}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var preparationTime int
	preparationTime = 2 * numberOfLayers
	//panic("PreparationTime not implemented")
	return preparationTime

}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var elapsedTime int
	elapsedTime = 2*numberOfLayers + actualMinutesInOven
	//panic("ElapsedTime not implemented")
	return elapsedTime

}

func main() {
	remainingOvenTime := RemainingOvenTime(10)
	fmt.Println("Remaining Oven Time:", remainingOvenTime)
	prepTime := PreparationTime(1)
	fmt.Println("Preparation Time:", prepTime)
	elapsedtime := ElapsedTime(3, 20)
	fmt.Println("Elapsed Time:", elapsedtime)
}
