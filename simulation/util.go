package simulation

import (
	"math/rand"
)

// MakeRange generates a sequence of int numbers
func MakeRange(min, max int) []int {
	res := make([]int, max - min)
	for i := range res {
		res[i] = min + i
	}
	return res
}

// Shuffle input int array using a random number generator
func Shuffle(vals []int, r *rand.Rand) {
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

// Sum sequence of integers
func Sum(input ...int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}
