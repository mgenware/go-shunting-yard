package main

import (
	"fmt"
)

func main() {
	// expression
	input := []string{"12", "+", "3", "*", "(", "4", "-", "54", ")", "^", "2", "-", "4", "/", "2", "*", "12"}
	// expression to Reverse Polish notation (RPN) tokens
	tokens, err := parse(input)
	if err != nil {
		panic(err)
	}
	for _, token := range tokens {
		fmt.Printf("%v(%v) ", token.Value, token.Type)
	}
	fmt.Println()

	// evaluate tokens
	val, err := evaluate(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
