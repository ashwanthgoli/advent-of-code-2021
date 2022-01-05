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

func lanternfish(input string, days int) (int, error) {
	input = strings.TrimSpace(input)
	fishes := make([]int, 0, 1000)

	for len(input) > 0 {
		m := strings.IndexByte(input, ',')
		var (
			err error
			val int
		)
		if m == -1 {
			val, err = strconv.Atoi(input)
			input = ""
		} else {
			val, err = strconv.Atoi(input[:m])
			input = input[m+1:]
		}

		if err != nil {
			return 0, err
		}
		fishes = append(fishes, val)
	}

	for i := 0; i < days; i++ {
		cnt := 0

		for j := range fishes {
			if fishes[j] == 0 {
				fishes[j] = 6
				cnt++
			} else {
				fishes[j]--
			}
		}

		for j := 0; j < cnt; j++ {
			fishes = append(fishes, 8)
		}

		// fmt.Printf("After day %d: %v\n", i+1, fishes)
	}

	return len(fishes), nil
}

func lanternfishV2(input string, days int) (int64, error) {
	var fishesOnDay [7]int64

	for i := 0; i < len(input); i = i + 2 {
		fishesOnDay[input[i]-'0']++
	}

	var incubatingD0, incubatingD1 int64
	index := 0
	for day := 0; day < days; day++ {
		incubated := incubatingD0
		incubatingD0 = incubatingD1
		incubatingD1 = fishesOnDay[index]

		fishesOnDay[index] += incubated

		if index == 6 {
			index = -1
		}
		index++
	}

	sum := incubatingD0 + incubatingD1
	for _, v := range fishesOnDay {
		sum += v
	}

	return sum, nil
}

func main() {
	resV1, err := lanternfish(input, 80)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result %d\n", resV1)

	resV2, err := lanternfishV2(input, 256)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result %d\n", resV2)
}
