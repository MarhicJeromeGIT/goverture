package goverture

import (
	"reflect"
	"sort"
)

// slicesEqual checks if two slices of ints are equal, ignoring order.
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sortedA := append([]int(nil), a...)
	sortedB := append([]int(nil), b...)
	sort.Ints(sortedA)
	sort.Ints(sortedB)
	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}
	return true
}

// findRowIndex returns the index of the row in the matrix that matches the given row slice.
func FindRowIndex(matrix [][]int, row []int) (int, bool) {
	for i, mrow := range matrix {
		if reflect.DeepEqual(mrow, row) {
			return i, true
		}
	}
	return -1, false
}
