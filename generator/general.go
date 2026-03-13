package generator

import (
	"encoding/json"
	"math/rand/v2"
	"os"
)

// A helper function for loading in data (i.e., JSON files) from
// ./data.  This function exists for the sake of reducing bloat
// (since reading in JSON files in Golang is, well, verbose);
// because of this, also make this a generic function that
// accepts a struct in generator/structures.go
func loadJson[T any](jsonPath string) (T, error) {
	var result T

	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

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
