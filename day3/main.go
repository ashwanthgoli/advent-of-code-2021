package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func decode(input string) (int64, error) {
	var cum []int
	input = strings.TrimSpace(input)

	for len(input) > 0 {
		m := strings.Index(input, "\n")

		var line string
		if m == -1 {
			line = input
			input = ""
		} else {
			line = input[:m]
			input = input[m+1:]
		}

		if cum == nil {
			cum = make([]int, len(line))
		}

		for j, c := range line {
			if c == '0' {
				cum[j] -= 1
			} else {
				cum[j] += 1
			}
		}
	}

	var gamma, epsilon int64
	for _, c := range cum {
		gamma <<= 1
		epsilon <<= 1

		if c > 0 {
			gamma |= 1
		} else if c < 0 {
			epsilon |= 1
		} else {
			return 0, fmt.Errorf("Not able to find most & least common bit")
		}
	}

	return gamma * epsilon, nil
}

//go:embed input.txt
var input string

func main() {
	res, err := decode(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result %d\n", res)
}
