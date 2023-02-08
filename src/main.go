package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (c calc) operate(in string, op string) int {
	values := strings.Split(in, op)
	value1 := parser(values[0])
	value2 := parser(values[1])

	switch op {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "*":
		return value1 * value2
	case "/":
		return value1 / value2
	default:
		fmt.Println(op, "- error")
		os.Exit(1)
		return 1
	}
}

func parser(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return value
}

func ReadIn() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

// func main() {
// 	in := readIn()
// 	op := readIn()

// 	c := calc{}
// 	fmt.Println(c.operate(in, op))
// }
