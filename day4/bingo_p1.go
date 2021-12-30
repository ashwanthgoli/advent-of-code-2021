package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

const bSize int = 5

type noValuesError struct{}

func (_ noValuesError) Error() string {
	return "No more values to return"
}

type bingoBoard [bSize][bSize]struct {
	elem   int
	marked bool
}

func (b *bingoBoard) crossNumber(num int) {
Outer:
	for i := 0; i < bSize; i++ {
		for j := 0; j < bSize; j++ {
			if b[i][j].elem == num {
				b[i][j].marked = true
				break Outer
			}
		}
	}
}

func (b *bingoBoard) sumUnmarked() int {
	sum := 0
	for i := 0; i < bSize; i++ {
		for j := 0; j < bSize; j++ {
			if !b[i][j].marked {
				sum += b[i][j].elem
			}
		}
	}

	return sum
}

func (b *bingoBoard) printBoard() {
	var buf bytes.Buffer
	buf.WriteString("\n")
	for i := 0; i < bSize; i++ {
		for j := 0; j < bSize; j++ {
			if b[i][j].marked {
				buf.WriteString("_")
			} else {
				buf.WriteString(fmt.Sprintf("%2d", b[i][j].elem))
			}
			buf.WriteString(" ")
		}
		buf.WriteString("\n")
	}

	fmt.Printf(buf.String())
}

func (b *bingoBoard) isComplete() bool {
	for i := 0; i < bSize; i++ {
		var c1, c2 int
		for j := 0; j < bSize; j++ {
			if b[i][j].marked {
				c1++
			}
			if b[j][i].marked {
				c2++
			}
		}

		if c1 == bSize || c2 == bSize {
			return true
		}
	}

	return false
}

func drawNumberGen(input string) func() (int, error) {
	m := strings.IndexByte(input, '\n')
	l1 := input[:m]

	return func() (int, error) {
		if len(l1) == 0 {
			return 0, noValuesError{}
		}

		n := strings.IndexByte(l1, ',')
		var (
			val int
			err error
		)

		if n == -1 {
			val, err = strconv.Atoi(l1)
			l1 = ""
		} else {
			val, err = strconv.Atoi(l1[:n])
			l1 = l1[n+1:]
		}

		if err != nil {
			return 0, fmt.Errorf("Invalid token, expected number")
		}

		return val, nil
	}
}

func skipLine(input string) string {
	m := strings.IndexByte(input, '\n')
	if m == -1 {
		return input
	}
	return input[m+1:]
}

func getBoards(input string) ([]bingoBoard, error) {
	var (
		n int = strings.Count(input, "\n") / 6
	)

	input = skipLine(input)

	boards := make([]bingoBoard, 0, n)
	for len(input) > 0 {
		input = skipLine(input)

		var board bingoBoard
		for i := 0; i < bSize; i++ {
			m := strings.IndexByte(input, '\n')
			line := input[:m]

			var j int
			for len(line) > 0 {
				s := strings.IndexFunc(line, func(r rune) bool { return !unicode.IsSpace(r) })

				if s != -1 {
					line = line[s:]
				}

				e := strings.IndexByte(line, ' ')
				var (
					elem int
					err  error
				)

				if e == -1 {
					elem, err = strconv.Atoi(line)
					line = ""
				} else {
					elem, err = strconv.Atoi(line[:e])
					line = line[e+1:]
				}

				if err != nil {
					return nil, errors.New("Unexpected format\n")
				}

				board[i][j].elem = elem
				j++
			}

			input = input[m+1:]
		}

		boards = append(boards, board)
	}

	return boards, nil
}

func main() {
	boards, err := getBoards(input)
	if err != nil {
		fmt.Printf("Failed to read boards input. error %s", err.Error())
		os.Exit(1)
	}

	var e error
	drawFn := drawNumberGen(input)
	for {
		var (
			v      int
			result int
		)
		v, e = drawFn()
		if e != nil {
			if _, ok := e.(noValuesError); ok {
				break
			} else {
				fmt.Printf("Failed to draw the next number. error: %s", e.Error())
				os.Exit(1)
			}
		}

		for i := range boards {
			boards[i].crossNumber(v)

			if boards[i].isComplete() {
				sum := boards[i].sumUnmarked()

				if sum*v > result {
					result = sum * v
				}
			}
		}

		if result > 0 {
			fmt.Printf("Result: %d\n", result)
			os.Exit(0)
		}
	}
}
