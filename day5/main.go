package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

const gridSize int = 1000

type Grid [gridSize][gridSize]int

//go:embed input.txt
var input string

func VentsP1(input string) (int, error) {
	var grid Grid

	for len(input) > 0 {
		m := strings.IndexByte(input, '\n')

		var (
			line           string
			x1, y1, x2, y2 int
		)

		if m == -1 {
			line = input
			input = ""
		} else {
			line = input[:m]
			input = input[m+1:]
		}

		items, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil && items != 4 {
			return 0, fmt.Errorf("Failed to parse line: %s", line)
		}

		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			for j := y1; j <= y2; j++ {
				grid[x1][j]++
			}

		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			for i := x1; i <= x2; i++ {
				grid[i][y1]++
			}
		}
	}

	res := 0
	for _, r := range grid {
		for _, e := range r {
			if e >= 2 {
				res++
			}
		}
	}

	return res, nil
}

func VentsP2(input string) (int, error) {
	var grid Grid

	for len(input) > 0 {
		m := strings.IndexByte(input, '\n')

		var (
			line           string
			x1, y1, x2, y2 int
		)

		if m == -1 {
			line = input
			input = ""
		} else {
			line = input[:m]
			input = input[m+1:]
		}

		items, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil && items != 4 {
			return 0, fmt.Errorf("Failed to parse line: %s", line)
		}

		switch {
		case x1 == x2:
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			for j := y1; j <= y2; j++ {
				grid[x1][j]++
			}
		case y1 == y2:
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			for i := x1; i <= x2; i++ {
				grid[i][y1]++
			}
		case x1+y1 == x2+y2:
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			for i, j := x1, y1; i <= x2; i, j = i+1, j-1 {
				grid[i][j]++
			}
		case x1-y1 == x2-y2:
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			for i, j := x1, y1; i <= x2; i, j = i+1, j+1 {
				grid[i][j]++
			}
		}
	}

	res := 0
	for _, r := range grid {
		for _, e := range r {
			if e >= 2 {
				res++
			}
		}
	}

	return res, nil
}

func main() {
	res, err := VentsP1(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", res)

	res, err = VentsP2(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", res)
}
