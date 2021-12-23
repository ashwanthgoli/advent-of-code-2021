package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func navigateP1(input string) (int64, error) {
	var (
		x, y int64
	)

	input = strings.TrimSpace(input)
	for i := 0; i < len(input); {
		var (
			step int64
			line string
		)

		e := strings.Index(input[i:], "\n")
		if e == -1 {
			line = input[i:]
			i = len(input)
		} else {
			line = input[i : i+e]
			i += e + 1
		}

		s := strings.Index(line, " ")
		if s == -1 {
			return 0, fmt.Errorf("line doesn't match the expected input format: %s", line)
		}

		step, err := strconv.ParseInt(line[s+1:], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("line doesn't match the expected input format: %s, err: %s", line, err.Error())
		}
		switch line[:s] {
		case "forward":
			x += step
		case "down":
			y += step
		case "up":
			y -= step
		}
	}

	return x * y, nil
}

func navigateP2(input string) (int64, error) {
	var (
		ypos, xpos, aim int64
	)
	for i := 0; i < len(input); {
		m := strings.Index(input[i:], "\n")

		var line string
		if m == -1 {
			line = input[i:]
			i = len(input)
		} else {
			line = input[i : i+m]
			i += m + 1
		}

		n := strings.Index(line, " ")
		if n == -1 {
			return 0, fmt.Errorf("Invalid input format. %s", line)
		}

		val, err := strconv.ParseInt(line[n+1:], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("Invalid input format. %s", line)
		}

		switch line[:n] {
		case "forward":
			xpos += val
			ypos += val * aim
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}

	return xpos * ypos, nil
}

func main() {
	out, err := navigateP1(input)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result P1 %d\n", out)

	out, err = navigateP2(input)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result P2 %d\n", out)
}
