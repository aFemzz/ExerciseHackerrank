package main

import (
	answer "ExerciseTest/handlers"
	"fmt"
)

func main() {
	quantity := []int{1, 2, 4}
	m := 4
	fmt.Println(answer.GetMaximumRevenue(quantity, m)) // output should be 11

}
