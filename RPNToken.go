package shuntingYard

const (
	TokenTypeOperand  = 1
	TokenTypeOperator = 2
)

// RPNToken represents a abstract token object in RPN(Reverse Polish notation) which could either be an operator or operand.
type RPNToken struct {
	Type  int
	Value string
}

// IsOperand returns whether a token is an operand with a specified value.
func (token *RPNToken) IsOperand(val string) {
	return token.Type == TokenTypeOperand && token.Value == val
}

// IsOperator returns whether a token is an operator with a specified value.
func (token *RPNToken) IsOperator(val string) {
	return token.Type == TokenTypeOperand && token.Value == val
}
