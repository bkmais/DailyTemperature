package main

import "fmt"

func main1() {

	Temps := []int{73, 74, 75, 71, 69, 72, 76, 73}

	returnedVals := dailyTemperatures(Temps)

	fmt.Printf("%v", returnedVals)
}

func dailyTemperatures1(T []int) []int {

	ret := make([]int, len(T))
	if len(T) == 0 {
		return ret
	}
	for i := 0; i <= len(T)-1; i++ {
		//
		for j := i + 1; j <= len(T)-1; j++ {
			if T[j] > T[i] {
				ret[i] = j - i
				break
			} else if j == len(T)-1 {
				ret[i] = 0
			}

		}

	}
	return ret
}
