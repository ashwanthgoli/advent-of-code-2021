package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func crabPosition(input string) (int, error) {
	input = strings.TrimSpace(input)
	positions := []int{}

	for len(input) > 0 {
		var (
			val int
			err error
		)

		if m := strings.IndexByte(input, ','); m == -1 {
			val, err = strconv.Atoi(input)
			input = ""
		} else {
			val, err = strconv.Atoi(input[:m])
			input = input[m+1:]
		}

		if err != nil {
			return 0, err
		}

		positions = append(positions, val)
	}

	sort.Ints(positions)
	median := positions[len(positions)/2]

	var fuel int
	for _, pos := range positions {
		if pos > median {
			fuel += pos - median
		} else {
			fuel += median - pos
		}
	}

	return fuel, nil
}

func crabPositionV2(input string) (int, error) {
	input = strings.TrimSpace(input)

	positions := []int{}
	maxPos := math.MinInt32
	minPos := math.MaxInt32
	for len(input) > 0 {
		var (
			val int
			err error
		)

		if m := strings.IndexByte(input, ','); m == -1 {
			val, err = strconv.Atoi(input)
			input = ""
		} else {
			val, err = strconv.Atoi(input[:m])
			input = input[m+1:]
		}

		if err != nil {
			return 0, err
		}

		if val > maxPos {
			maxPos = val
		}

		if val < minPos {
			minPos = val
		}

		positions = append(positions, val)
	}

	var fuel = math.MaxInt32
	for pos := minPos; pos <= maxPos; pos++ {
		var curFuel int
		for _, curPos := range positions {
			distance := curPos - pos
			if curPos < pos {
				distance = pos - curPos
			}

			curFuel += (distance + 1) * distance / 2
		}

		if curFuel < fuel {
			fuel = curFuel
		}
	}

	return fuel, nil
}

func main() {
	res, err := crabPosition(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result v1: %d\n", res)

	res, err = crabPositionV2(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result v2: %d\n", res)
}
