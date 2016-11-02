package shuntingYard

import (
	"errors"
	"math"
	"strconv"
)

// Evaluate evaluates a list of RPNTokens and returns calculated value.
func Evaluate(tokens []*RPNToken) (int, error) {
	if tokens == nil {
		return 0, errors.New("tokens cannot be nil")
	}

	var stack []int
	for _, token := range tokens {
		// push all operands to the stack
		if token.Type == TokenTypeOperand {
			val, err := strconv.Atoi(token.Value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, val)
		} else {
			// execute current operator
			if len(stack) < 2 {
				return 0, errors.New("Missing operand")
			}
			// pop 2 elements
			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			val, err := evaluateOperator(token.Value, arg1, arg2)
			if err != nil {
				return 0, err
			}
			// push result back to stack
			stack = append(stack, val)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("Stack corrupted")
	}
	return stack[len(stack)-1], nil
}

// executes an operator
func evaluateOperator(oper string, a, b int) (int, error) {
	switch oper {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "^":
		return int(math.Pow(float64(a), float64(b))), nil
	default:
		return 0, errors.New("Unknown operator: " + oper)
	}
}
