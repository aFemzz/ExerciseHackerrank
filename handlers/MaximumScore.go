package handlers

import "fmt"

func MaxAchievableScore(arr []int, moves int) int {
	if len(arr) == 0 || moves <= 0 {
		return 0
	}

	// making sure the array is in  accordance to the input
	if moves+1 < len(arr) {
		arr = arr[:moves+1]
	}
	// fmt.Println("INI  ARR : ", arr)

	totalScore := 0

	for i := 0; i < moves; i++ {

		maxIndex := findMaxIndex(arr)

		// totaling the score of maximum valued index

		totalScore += arr[maxIndex]

		arr = removeElementByIndex(arr, maxIndex)
	}

	return totalScore
}

// Find the index of the maximum value
func findMaxIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

// Remove the maximum value from the array
func removeElementByIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("index exceeded array length or 0")
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}
