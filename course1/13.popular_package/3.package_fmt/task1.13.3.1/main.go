package main

import (
	"fmt"
	"strings"
)

func generateMathString(operands []int, operator string) string {
	if len(operands) < 2 {
		return "not enough operands"
	}
	switch operator {
	case "+":
		return summ(operands)
	case "-":
		return sub(operands)
	case "/":
		return div(operands)
	case "*":
		return mult(operands)
	default:
		return "unknown operator"
	}
}

func summ(operands []int) string {
	res := 0
	str := make([]string, 0, len(operands))
	for _, operand := range operands {
		res += operand
		str = append(str, fmt.Sprintf("%d", operand))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " + "), res)
}

func sub(operands []int) string {
	res := operands[0]
	str := make([]string, 0, len(operands))
	str = append(str, fmt.Sprintf("%d", operands[0]))
	for i := 1; i < len(operands); i++ {
		res -= operands[i]
		str = append(str, fmt.Sprintf("%d", operands[i]))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " - "), res)
}

func mult(operands []int) string {
	res := 1
	str := make([]string, 0, len(operands))
	for _, operand := range operands {
		res *= operand
		str = append(str, fmt.Sprintf("%d", operand))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " * "), res)
}

func div(operands []int) string {
	res := float64(operands[0])
	str := make([]string, 0, len(operands))
	str = append(str, fmt.Sprintf("%d", operands[0]))
	for i := 1; i < len(operands); i++ {
		res /= float64(operands[i])
		str = append(str, fmt.Sprintf("%d", operands[i]))
	}
	return fmt.Sprintf("%s = %.2f", strings.Join(str, " / "), res)
}
