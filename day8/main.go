package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func sevenSegCnt(input string) (int, error) {
	var cnt int
	for {
		m := strings.IndexByte(input, '\n')
		if m == -1 {
			break
		}

		line := input[:m]
		input = input[m+1:]

		delim := strings.IndexByte(line, '|')
		if delim == -1 {
			return 0, errors.New("Invalid input. Missing delimiter")
		}

		desiredNums := map[int]struct{}{
			2: {}, // 1
			3: {}, // 7
			4: {}, // 4
			7: {}, // 8
		}
		digits := line[delim+2:]

		fmt.Println(digits)
		for {
			if n := strings.IndexByte(digits, ' '); n == -1 {
				fmt.Printf("%d\n", len(digits))
				if _, ok := desiredNums[len(digits)]; ok {
					cnt++
				}
				break
			} else {
				fmt.Printf("%d ", n)
				if _, ok := desiredNums[n]; ok {
					cnt++
				}
				digits = digits[n+1:]
			}
		}
		fmt.Printf("Cnt %d\n", cnt)
	}

	return cnt, nil
}

//  aaaa
// b    c
// b    c
// b    c
//  dddd
// e    f
// e    f
// e    f
// e    f
//  gggg
//
//     abcdefg
// 1 - 0010010
// 2 - 1011101
// 3 - 1011011
// 4 - 0111010
// 5 - 1101011
// 6 - 1101111
// 7 - 1010010
// 8 - 1111111
// 9 - 1111011
func sevenSegRead() {
	_ = map[int]int{
		1: 0010010,
		2: 1011101,
		3: 1011011,
		4: 0111010,
		5: 1101011,
		6: 1101111,
		7: 1010010,
		8: 1111111,
		9: 1111011,
	}
}

func main() {
	res, err := sevenSegCnt(input)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result p1: %d\n", res)
}
