package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

func decodePC(input string) (int64, error) {
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

func findMCB(lines []string, j int) byte {
	var on, off int

	for _, line := range lines {
		if line[j] == '1' {
			on += 1
		} else {
			off += 1
		}
	}

	var msb byte
	if on > off {
		msb = '1'
	} else if off > on {
		msb = '0'
	} else {
		msb = '-'
	}

	return msb
}

func binToInt(input string) int64 {
	var v int64

	for _, c := range input {
		v <<= 1

		if c == '1' {
			v |= 1
		}
	}

	return v
}

func decodeLSR_V1(input string) (int64, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	if len(lines) == 0 {
		return 0, nil
	}

	o2set := lines
	for j := 0; len(o2set) > 1; j++ {
		msb := findMCB(o2set, j)

		var tmp []string
		for _, line := range o2set {
			if line[j] == msb || (msb == '-' && line[j] == '1') {
				tmp = append(tmp, line)
			}
		}

		o2set = tmp
	}

	co2set := lines
	for j := 0; len(co2set) > 1; j++ {
		msb := findMCB(co2set, j)

		var tmp []string
		for _, line := range co2set {
			if (msb == '-' && line[j] == '0') || (msb != '-' && line[j] != msb) {
				tmp = append(tmp, line)
			}
		}

		co2set = tmp
	}

	if len(o2set) != 1 {
		return 0, fmt.Errorf("Failed to filter down to a single entry while decoding o2 rating")
	}
	o2r := binToInt(o2set[0])

	if len(co2set) != 1 {
		return 0, fmt.Errorf("Failed to filter down to a single entry while decoding co2 rating")
	}
	co2r := binToInt(co2set[0])

	return o2r * co2r, nil
}

func parseReport(n, blen int, input string, mc bool) int64 {
	filtered := make([]int, 0, n)
	for i := 0; i < n; i++ {
		filtered = append(filtered, i)
	}

	bit := 0
	for len(filtered) > 1 {
		var ones, zeroes int
		for _, r := range filtered {
			if input[r*(blen+1)+bit] == '1' {
				ones++
			} else {
				zeroes++
			}
		}

		var filter byte = '1'
		if mc {
			filter = '1'
			if zeroes > ones {
				filter = '0'
			}
		} else {
			filter = '0'
			if ones < zeroes {
				filter = '1'
			}
		}

		var j int
		for _, r := range filtered {
			if input[r*(blen+1)+bit] == filter {
				filtered[j] = r
				j++
			}
		}
		filtered = filtered[:j]

		bit++
	}

	s := filtered[0] * (blen + 1)
	return binToInt(input[s : s+blen])
}

func decodeLSR_V2(input string) (int64, error) {
	var (
		blen int = strings.IndexByte(input, '\n')
		n    int = len(input) / (blen + 1)
	)

	return parseReport(n, blen, input, true) * parseReport(n, blen, input, false), nil
}

//go:embed input.txt
var input string

func main() {
	res, err := decodePC(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Power consumption: %d\n", res)

	res, err = decodeLSR_V1(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Life support rating: %d\n", res)

	res, err = decodeLSR_V2(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Life support rating: %d\n", res)
}
