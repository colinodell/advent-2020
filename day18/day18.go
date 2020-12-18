package main

import (
	"advent-2020/utils"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	input := utils.ReadLines("./day18/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Sum of all expressions: %d\n", EvaluateSum(input, Part1))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Sum of all expressions: %d\n", EvaluateSum(input, Part2))
}

type TokenType int
const (
	Operand TokenType = iota
	Operator
)

type Token struct {
	Type TokenType
	Value interface{}
}

type OperatorPrecedence map[string]int

var Part1 = OperatorPrecedence {"+": 0, "*": 0}
var Part2 = OperatorPrecedence {"+": 1, "*": 0}

func Evaluate(expression string, precedence OperatorPrecedence) int {
	return evaluate(parse(scan(expression), precedence))
}

func EvaluateSum(expressions []string, precedence OperatorPrecedence) int {
	sum := 0
	for _, line := range expressions {
		sum += Evaluate(line, precedence)
	}

	return sum
}

// Splits the string into its individual components
func scan(input string) []string {
	var result []string

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		result = append(result, s.TokenText())
	}

	return result
}

// Implements Dijkstra's shunting-yard algorithm to parse the expression into RPN tokens
func parse(tokens []string, operators map[string]int) []Token {
	var result []Token
	var operatorStack []string

	for _, token := range tokens {
		// Is this token an operand (aka number)?
		if num, err := strconv.Atoi(token); err == nil {
			result = append(result, Token{Type: Operand, Value: num})
			continue
		}

		if token == "(" {
			operatorStack = append(operatorStack, token)
			continue
		}

		if token == ")" {
			// Pop operators off the stack until we find a matching open paren
			for len(operatorStack) > 0 {
				op := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if op == "(" {
					break
				}

				result = append(result, Token{Type: Operator, Value: op})
			}
			continue
		}

		// We have an operator
		// Determine its priority
		currentTokenPriority, ok := operators[token]
		if !ok {
			panic("unknown operator: " + token)
		}

		for len(operatorStack) > 0 {
			// Check top of the operator stack
			op := operatorStack[len(operatorStack)-1]
			if op == "(" {
				break
			}

			poppedOpPriority := operators[op]
			if currentTokenPriority <= poppedOpPriority {
				// Pop it off the stack and onto our result token set
				operatorStack = operatorStack[:len(operatorStack)-1]
				result = append(result, Token{Type: Operator, Value: op})
			} else {
				break
			}
		}

		operatorStack = append(operatorStack, token)
	}

	// Handle any remaining operators
	for len(operatorStack) > 0 {
		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		result = append(result, Token{Type: Operator, Value: op})
	}

	return result
}

func evaluate(tokens []Token) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		// Push operands (numbers) onto the stack as we find them
		if token.Type == Operand {
			stack = append(stack, token.Value.(int))
			continue
		}

		// Pop two operands off the stack and apply them
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		// Perform the operation and push the result back onto the stack
		stack = append(stack, doOperation(token.Value.(string), a, b))
	}

	// The result will be the last item in the stack
	return stack[0]
}

func doOperation(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	}

	panic("unknown operator: " + op)
}
