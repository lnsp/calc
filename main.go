package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lnsp/rechner/lib"
)

var OperatorError = errors.New("Operator error")
var ParseError = errors.New("Unexpected symbol")

func isNumber(token string) bool {
	isnumber, err := regexp.Match("[0-9]+", []byte(token))
	if err != nil {
		panic(err)
	}
	return isnumber
}

func isOperator(token string) bool {
	for _, s := range []string{"+", "-", "*", "/"} {
		if token == s {
			return true
		}
	}
	return false
}

func getPrecedence(operator string) int {
	switch operator {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	}
	return 0
}

func performOperation(operator string, a, b int) (int, error) {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	return result, nil
}

func parse(line string) (lib.OperatorQueue, error) {
	stack := lib.NewStack()
	queue := lib.NewQueue()

	// Scan tokens
	tokens := strings.Split(line, " ")
	for i := 0; i < len(tokens); i++ {
		item := strings.TrimSpace(tokens[i])
		if isNumber(item) {
			queue.Append(item)
		} else if isOperator(item) {
			precedence := getPrecedence(item)
			for !stack.IsEmpty() && isOperator(stack.Peek()) && getPrecedence(stack.Peek()) >= precedence {
				queue.Append(stack.Pop())
			}
			stack.Push(item)
		} else if item == "(" {
			stack.Push(item)
		} else if item == ")" {
			for !stack.IsEmpty() && stack.Peek() != "(" {
				queue.Append(stack.Pop())
			}
			stack.Pop()
		} else if item != "" {
			return queue, ParseError
		}
	}

	for !stack.IsEmpty() {
		queue.Append(stack.Pop())
	}

	return queue, nil
}

func eval(rpn lib.OperatorQueue) (int, error) {
	stack := lib.NewIntStack()

	for !rpn.IsEmpty() {
		op := rpn.Poll()
		if isOperator(op) {
			b, a := stack.Pop(), stack.Pop()
			result, err := performOperation(op, a, b)
			if err != nil {
				return result, err
			}
			stack.Push(result)
		} else {
			val, _ := strconv.Atoi(op)
			stack.Push(val)
		}
	}

	return stack.Pop(), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()

		rpn, err := parse(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		result, err := eval(rpn)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)
		fmt.Print("> ")
	}
}
