# Go Cover Solver

A Go library for solving exact cover problems using Donald Knuth's Algorithm X with the Dancing Links technique (DLX).

## Features

- Solve exact cover problems efficiently.
- Simple API to integrate into your Go projects.
- Solutions are streamed via a Go channel for flexibility in handling results.

## Installation

To install the library, run the following command:

```bash
go get github.com/MarhicJeromeGIT/goverture
```

## Usage

You can use this library to solve exact cover problems by passing a 2D matrix (slice of slices) to the `SolveDLX` function. The function returns a channel that streams solutions.

### Example

```go
package main

import (
    "context"
    "fmt"
    "github.com/MarhicJeromeGIT/goverture"
)

func main() {
    matrix := [][]int{
        {1, 0, 0, 1, 0, 0, 1},
        {1, 0, 0, 1, 0, 0, 0},
        {0, 0, 0, 1, 1, 0, 1},
        {0, 1, 1, 0, 0, 1, 0},
        {0, 1, 0, 0, 0, 0, 1},
        {0, 0, 1, 1, 1, 0, 0},
    }

    ctx := context.Background()
    solutions := cover_solver.SolveDLX(ctx, matrix)

    for solution := range solutions {
        fmt.Println("Solution:", solution)
    }
}
```

### Function Signature

```go
// SolveDLX initiates the DLX search and returns a channel of solutions
func SolveDLX(ctx context.Context, matrix [][]int) <-chan [][]int
```

- **`ctx context.Context`**: Used for cancellation and deadlines.
- **`matrix [][]int`**: The matrix representing the exact cover problem.
- **Returns**: A channel that streams all valid solutions (`[][]int`).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

