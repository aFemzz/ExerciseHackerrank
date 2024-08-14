package main

import (
	answer "ExerciseTest/handlers"
	"fmt"
)

func main() {
	// TASK 4 - Maximize The Revenue
	quantity := []int{1, 2, 4}
	m := 4
	fmt.Println("Max revenue for the store: ", answer.GetMaximumRevenue(quantity, m)) // output should be 11

	// TASK 3 - Maximum Score 2
	arr := []int{4, 6, -10, -1, 10, -20}
	moves := 4
	result := answer.MaxAchievableScore(arr, moves)
	fmt.Println("The maximum achievable score is:", result)

}
