package main

import (
	"fmt"
)

func main() {
	input := []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"}
	//input := []string{"(", "1", "+", "2", ")", "^", "3"}
	tokens, err := parse(input)
	if err != nil {
		panic(err)
	}
	for _, token := range tokens {
		fmt.Printf("%v(%v) ", token.Value, token.Type)
	}
}
