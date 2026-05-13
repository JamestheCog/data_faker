package generator

import (
	"math/rand/v2"
)

// A helper function that, given a slice and a number "n", generate
// a subset of "n" items from the said slice.  Duplicates are
// purposely allowed here:
func randomSample[T any](slice []T, n int) []T {
	toReturn := make([]T, n)

	for i := 0; i < len(toReturn); i++ {
		j := rand.IntN(len(slice))
		toReturn[i] = slice[j]
	}
	return toReturn
}

// A generic function that essentially removes an element from a slice.
// 'nuff said.
func removeElement[T comparable](slice []T, element T) []T {
	splitInd := -1
	for i, v := range slice {
		if v == element {
			splitInd = i
			break
		}
	}

	if splitInd < 0 {
		return slice
	}
	return append(slice[:splitInd], slice[(splitInd+1):]...)
}
