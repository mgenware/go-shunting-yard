package shuntingYard

import "testing"

func TestEvaluate(t *testing.T) {
	tokens, err := Parse([]string{"1", "-", "32", "*", "3"})
	if err != nil {
		panic(err)
	}

	var expected, got int
	expected = -95
	got, err = Evaluate(tokens)
	if err != nil {
		panic(err)
	}
	if got != expected {
		t.Fatalf("Expected %v, Got %v.", expected, got)
	}

	tokens, err = Parse([]string{"1", "+", "2", "-", "3", "*", "4", "/", "5"})
	if err != nil {
		panic(err)
	}

	expected = 1
	got, err = Evaluate(tokens)
	if err != nil {
		panic(err)
	}
	if got != expected {
		t.Fatalf("Expected %v, Got %v.", expected, got)
	}
}
