package predictor

import (
	"fmt"
	"math/rand"
	"time"
)

func predictMarkov(arr []int, numStates int, predictionWindow int) (res []int) {

	// states := 10

	//matrixSize := states+1

	//x := []int{1,2,3,1,2,1,2,1,2,1,3,3,3,3,3,1}
	// x := []int{6, 5, 4, 4, 4, 3, 4, 4, 3, 4, 3, 5, 3, 4, 4, 5, 3, 3, 5, 4, 3, 3, 5, 7, 4, 5, 5, 4}

	n := len(arr)

	fmt.Print(arr, "\n")
	fmt.Print(n, "\n")

	var p [][]float32
	for i := 0; i <= numStates; i++ {
		p = append(p, make([]float32, numStates+1))
	}

	for t := 0; t < n-1; t++ {
		f := arr[t]
		s := arr[t+1]
		p[f][s] = p[f][s] + 1
	}

	for i := 1; i <= numStates; i++ {
		sum := sumOfAllElements(p[i])
		for j := 1; j <= numStates; j++ {
			if sum == 0 {
				p[i][j] = 0
			} else {
				p[i][j] = p[i][j] / sum
			}
		}

	}
	results := make([]int, predictionWindow)
	lastElement := arr[len(arr)-1]
	for i := 0; i < predictionWindow; i++ {
		results[i] = predictNext(lastElement, p, predictionWindow)

		lastElement = results[i]
	}
	// fmt.Print(p, "\n")
	return results

}

func predictNext(lastElement int, transitionMatrix [][]float32, D int) int {
	var maxIndices = make([]int, D)
	max := transitionMatrix[lastElement][0]
	for _, elem := range transitionMatrix[lastElement] {
		if elem > max {
			max = elem
		}
	}
	/// Special case all zeros
	if max == 0 {
		return random(1, len(transitionMatrix[0]))
	}

	index := 0
	for i, elem := range transitionMatrix[lastElement] {
		if elem == max {
			maxIndices[index] = i
			index++
		}
	}

	return maxIndices[random(0, len(maxIndices))]

}

// func findMaxProbabilityState(int rowNumber, transitionMatrix [][]float32) int {
//
// }

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func sumOfAllElements(array []float32) (sum float32) {
	for _, i := range array {
		sum += i
	}
	//fmt.Print("Sum",sum)
	return
}
