package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func divep1v1(input string) (int, error) {
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

func divep1v2(input string) (int, error) {
	input = strings.TrimSpace(input)

	var (
		cnt, i int
		p      struct {
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

func divep2v1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nums := make([]int64, len(lines))

	for i, line := range lines {
		e, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("expected number. found token %s", line)
		}

		nums[i] = e
	}

	var (
		cnt  int
		prev struct {
			val int64
			set bool
		}
	)

	for i := 0; i+2 < len(nums); i++ {
		cur := nums[i] + nums[i+1] + nums[i+2]

		if !prev.set {
			prev.set = true
		} else if cur > prev.val {
			cnt++
		}

		prev.val = cur
	}

	return cnt, nil
}

func divep2v2(input string) (int, error) {
	input = strings.TrimSpace(input)

	var (
		cnt  int
		j    int
		prev [3]int64
	)

	for i := 0; i < len(input); {
		m := strings.Index(input[i:], "\n")

		var (
			cur int64
			err error
		)

		if m == -1 {
			cur, err = strconv.ParseInt(input[i:], 10, 64)
			i = len(input)
		} else {
			cur, err = strconv.ParseInt(input[i:i+m], 10, 64)
			i += m + 1
		}

		if err != nil {
			return 0, fmt.Errorf("expected number, found invalid token starting at index %d", i)
		}

		if j != 3 {
			prev[2-j] = cur
			j++
		} else {
			if cur+prev[0]+prev[1] > prev[0]+prev[1]+prev[2] {
				cnt++
			}

			prev[0], prev[1], prev[2] = cur, prev[0], prev[1]
		}
	}

	return cnt, nil
}

func main() {
	cnt, err := divep1v1(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("P1V1 Answer %d\n", cnt)

	cnt, err = divep1v2(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("P1V2 Answer %d\n", cnt)

	cnt, err = divep2v1(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("P2V1 Answer %d\n", cnt)

	cnt, err = divep2v2(input)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("P2V2 Answer %d\n", cnt)
}
