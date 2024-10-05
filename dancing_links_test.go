package goverture

import (
	"context"
	"testing"
)

func TestBuildDLX(t *testing.T) {
	// Example usage
	matrix := [][]int{
		{1, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1},
		{0, 1, 0, 0, 0, 0, 1},
	}

	secondaryColumns := make(map[int]bool)
	for i := 0; i < len(matrix[0]); i++ {
		secondaryColumns[i] = false
	}

	res := BuildDLX(matrix, secondaryColumns)
	_ = res
	println("ok")
}

func TestSolveDLX(t *testing.T) {
	matrix := [][]int{
		{1, 0, 0, 1, 0, 0, 0}, // Row 0
		{0, 0, 0, 1, 1, 0, 1}, // Row 1
		{0, 0, 1, 0, 1, 1, 0}, // Row 2
		{0, 1, 1, 0, 0, 1, 1}, // Row 3
		{0, 1, 0, 0, 0, 0, 1}, // Row 4
		{1, 1, 1, 0, 0, 1, 0}, // Row 5
		{1, 1, 1, 1, 1, 1, 1}, // Row 6
	}

	solutionsChan := SolveDLX(context.Background(), matrix)

	// Collect all solutions into a slice
	var solutions [][]int
	for sol := range solutionsChan {
		var solutionIndices []int
		for _, row := range sol {
			index, found := FindRowIndex(matrix, row)
			if !found {
				t.Errorf("Row %v not found in the matrix", row)
				continue
			}
			solutionIndices = append(solutionIndices, index)
		}
		solutions = append(solutions, solutionIndices)
	}

	// Define expected solutions as slices of row indices
	expectedSolutions := [][]int{
		{0, 2, 4}, // Solution 1
		{1, 5},    // Solution 2
		{6},       // Solution 3
	}

	// Check if the number of solutions matches
	if len(solutions) != len(expectedSolutions) {
		t.Errorf("Expected %d solutions, got %d", len(expectedSolutions), len(solutions))
	}

	// Create a copy of expectedSolutions to track which have been found
	expectedFound := make([]bool, len(expectedSolutions))

	// Iterate through each collected solution
	for _, sol := range solutions {
		matched := false
		for i, expectedSol := range expectedSolutions {
			if !expectedFound[i] && slicesEqual(sol, expectedSol) {
				expectedFound[i] = true
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("Unexpected solution found: %v", sol)
		}
	}

	// Check if all expected solutions were found
	for i, found := range expectedFound {
		if !found {
			t.Errorf("Expected solution %v not found", expectedSolutions[i])
		}
	}
}
