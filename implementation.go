package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// СalculatePrefix обчислює результат префіксного виразу.
// Вхідний вираз має бути в форматі <оператор> <операнд1> <операнд2>.
// Приклад: "+ 3 4" поверне 7.
func СalculatePrefix(expression string) (int, error) {
	stack := []int{}
	tokens := strings.Fields(expression)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperand(token) {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			// Token is an operator
			if len(stack) < 2 {
				return 0, fmt.Errorf("Invalid expression: not enough operands for operator %s", token)
			}
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result, _ := performOperation(op1, op2, token)
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Invalid expression: too many operands")
	}

	return stack[0], nil
}

// isOperand перевіряє, чи є токен операндом.
func isOperand(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

// performOperation виконує обчислення між двома операндами за допомогою оператора.
func performOperation(op1, op2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		return op1 / op2, nil
	case "^":
		return power(op1, op2), nil
	default:
		return 0, fmt.Errorf("Invalid operator")
	}
}

// power виконує піднесення першого операнду до степеня другого операнду.
func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
