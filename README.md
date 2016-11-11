# go-shunting-yard
A simple implementation of shunting-yard algorithm in Go(Golang).

## Installation
```sh
go get github.com/mgenware/go-shunting-yard
```

## Example
```go
package main

import (
	"fmt"

	"github.com/mgenware/go-shunting-yard"
)

func main() {
	// calculate 12+4*3/5-2
	// parse infix notation to postfix notation(RPN)
	tokens, err := shuntingYard.Parse([]string{"12", "+", "4", "*", "3", "/", "5", "-", "2"})
	if err != nil {
		panic(err)
	}

	// evaluate RPN tokens
	result, err := shuntingYard.Evaluate(tokens)
	if err != nil {
		panic(err)
	}

	// print the result
	fmt.Printf("Result: %v", result)
}

```
