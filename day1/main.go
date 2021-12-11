package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func DiveV1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var (
		cnt       int
		prev      int64
		prevValid bool
	)

	for _, line := range lines {
		cur, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("expected number, found %s", line)
		}

		if !prevValid {
			prevValid = true
		} else if cur > prev {
			cnt++
		}

		prev = cur
	}

	return cnt, nil
}

func DiveV2(input string) (int, error) {
	input = strings.TrimSpace(input)

	var (
		cnt int
		i   = 0
		p   struct {
			val int64
			set bool
		}
	)

	for i < len(input) {
		var (
			cur int64
			err error
		)

		m := strings.Index(input[i:], "\n")
		if m == -1 {
			cur, err = strconv.ParseInt(input[i:], 10, 64)
			i = len(input)
		} else {
			cur, err = strconv.ParseInt(input[i:i+m], 10, 64)
			i += m + 1
		}

		if err != nil {
			return 0, fmt.Errorf("Expected number, found invalid token beginning at index %d", i)
		}

		if !p.set {
			p.set = true
		} else if cur > p.val {
			cnt++
		}

		p.val = cur
	}

	return cnt, nil
}

func main() {
	cnt, err := DiveV1(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("V1 Answer %d\n", cnt)

	cnt, err = DiveV2(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("V2 Answer %d\n", cnt)
}
