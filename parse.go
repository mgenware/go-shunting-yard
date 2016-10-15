package shuntingYard

import (
	"errors"
	"fmt"
	"strconv"
)

// precedence of operators
var priorities map[string]int

// associativities of operators
var associativities map[string]bool

func init() {
	priorities = make(map[string]int, 0)
	associativities = make(map[string]bool, 0)

	priorities["+"] = 0
	priorities["-"] = 0
	priorities["*"] = 1
	priorities["/"] = 1
	priorities["^"] = 2

	// if not set, associativity will be false(left-associative)
	associativities["^"] = true
}

// Parse parses an array of token strings and returns an array of abstract tokens
// using Shunting-yard algorithm.
func Parse(tokens []string) ([]*RPNToken, error) {
	var ret []*RPNToken

	var operators []string
	for _, token := range tokens {
		if isOperand(token) {
			ret = append(ret, &RPNToken{Type: TokenTypeOperand, Value: token})
		} else {
			// check parentheses
			if token == "(" {
				operators = append(operators, token)
			} else if token == ")" {
				foundLeftParenthesis := false
				// pop until "(" is fouund
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if oper == "(" {
						foundLeftParenthesis = true
						break
					} else {
						ret = append(ret, &RPNToken{Type: TokenTypeOperator, Value: oper})
					}
				}
				if !foundLeftParenthesis {
					return nil, errors.New("Mismatched parentheses found")
				}
			} else {
				// operator priority and associativity
				priority, ok := priorities[token]
				if !ok {
					return nil, fmt.Errorf("Unknown operator: %v", token)
				}
				rightAssociative := associativities[token]

				for len(operators) > 0 {
					top := operators[len(operators)-1]

					if top == "(" {
						break
					}

					prevPriority := priorities[top]

					if (rightAssociative && priority < prevPriority) || (!rightAssociative && priority <= prevPriority) {
						// pop current operator
						operators = operators[:len(operators)-1]
						ret = append(ret, &RPNToken{Type: TokenTypeOperator, Value: top})
					} else {
						break
					}
				} // end of for len(operators) > 0

				operators = append(operators, token)
			} // end of if token == "("
		} // end of if isOperand(token)
	} // end of for _, token := range tokens

	// process remaining operators
	for len(operators) > 0 {
		// pop
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		if operator == "(" {
			return nil, errors.New("Mismatched parentheses found")
		}
		ret = append(ret, &RPNToken{Type: TokenTypeOperator, Value: operator})
	}
	return ret, nil
}

// isOperand determines whether a given string is an operand.
func isOperand(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
