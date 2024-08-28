package main

import (
	"exact_cover/set"
	"fmt"
	"sync"
)

const N = 5;

// Whether the given rows indexes form an exact cover of the matrix
func solved(matrix [][N]int, rows_index *set.Set) bool {
  columnCount := make([]int, N)

  for _, row_index := range rows_index.Elements() {
    row := matrix[row_index]
    for i, val := range row {
      if val == 1 {
      columnCount[i]++
      }
    }
  }

  // each column must be exactly 1
  for _, count := range columnCount {
    if count != 1 {
      return false
    }
  }
  return true
} 

// bruteforce backtracking to solve exact cover
func solve(matrix [][N]int) <-chan *set.Set {
  solution := set.NewSet()
  solutions := make(chan *set.Set, 10)

  wg := sync.WaitGroup{}
  wg.Add(1)
  go solveRec(matrix, solutions, solution, 0, &wg)

  wg.Wait()
  close(solutions)

  return solutions
}

// Simple backtracking solver
func solveRec(matrix [][N]int, solutions chan<- *set.Set, current_solution  *set.Set, min_index int, wg *sync.WaitGroup) {
  if min_index == 0 {
    // this is the top level/original call
    defer wg.Done()
  }

  if solved(matrix, current_solution) {
    fmt.Printf("Found solution %v\n", current_solution)
    solutions <- current_solution.Clone()
	  return
  }

  for i := min_index; i< len(matrix); i++ {
    if !current_solution.Contains(i) {
      // fmt.Println("Trying row ", i)
      current_solution.Add(i)
      solveRec(matrix, solutions, current_solution, i+1, wg)
      // fmt.Println("Backtrack from row ", i)
      current_solution.Remove(i)
    }
  }
}

func main() {
  matrix := make([][N]int, 0)
  // add some rows
  matrix = append(matrix, [N]int{1, 0, 0, 1, 0})
  matrix = append(matrix, [N]int{1, 1, 0, 1, 1})
  matrix = append(matrix, [N]int{1, 1, 1, 1, 1})
  matrix = append(matrix, [N]int{0, 0, 1, 0, 0})
  matrix = append(matrix, [N]int{1, 1, 1, 1, 0})
  matrix = append(matrix, [N]int{0, 0, 0, 0, 1})

  for _, row := range matrix {
	  fmt.Println(row)
  }

  // solve for exact cover by backtracking
  solutions := solve(matrix)
  // show all solutions
  for sol := range solutions {
    fmt.Println("solution", sol)
    fmt.Print("solved ?", solved(matrix, sol))
  }
}