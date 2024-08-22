package main

import "testing"

func TestFactorial(t *testing.T) {
	res1 := Factorial(0)
	if res1 != 1 {
		t.Errorf("Fatorial(0) = %d; want 1", res1)
	}

	res2 := Factorial(1)
	if res2 != 1 {
		t.Errorf("Fatorial(1) = %d; want 1", res2)
	}

	res3 := Factorial(5)
	if res3 != 120 {
		t.Errorf("Fatorial(5) = %d; want 120", res3)
	}

}
