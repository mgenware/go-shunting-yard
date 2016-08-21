package main

const (
	TokenTypeOperand  = 1
	TokenTypeOperator = 2
)

type Token struct {
	Type  int
	Value string
}
