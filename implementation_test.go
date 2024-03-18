package lab2

import (
	"fmt"
	_ "gopkg.in/check.v1"
	"testing"
)

// TestCalculatePrefix тест методу який обчислює результат префіксного виразу.
// Він перевіряє різні види префіксних виразів
// та перевіряє викид вийняткових ситуацій
func TestCalculatePrefix(t *testing.T) {
	tests := []struct {
		expression string
		expected   int
	}{
		// Simple expressions
		{"+ 3 4", 7},
		{"* + 3 4 5", 35},
		{"- * 7 4 2", 26},
		{"/ * 10 2 5", 4},
		{"+ - 15 3 2", 14},
		{"^ 2 3", 8}, // Test case for power operation
		// Complex expressions
		{"+ + + + + + 1 2 3 4 5 6 7", 28},
		{"* * * * * * 1 2 3 4 5 6 7", 5040},
		{"- - - - - - 1 2 3 4 5 6 7", -26},
		{"/ / / 100 5 2 2", 5},
		{"+ + + + + + + + + 1 2 3 4 5 6 7 8 9 10", 55},
		// Invalid expressions
		{"", 0},            // Empty expression
		{"^ 2 3 4", 0},     // Too many operands for power operation
		{"^ a 3", 0},       // Invalid operand
		{"& 1 2", 0},       // Invalid operator
		{"+ 1 2 3 4 x", 0}, // Invalid character
	}

	for _, test := range tests {
		result, err := calculatePrefix(test.expression)
		if err != nil {
			if test.expected != 0 {
				t.Errorf("Unexpected error for expression %s: %v", test.expression, err)
			}
		} else if result != test.expected {
			t.Errorf("Expected result %d, but got %d for expression %s", test.expected, result, test.expression)
		}
	}
}

func ExampleCalculatePrefix() {
	res, _ := calculatePrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 4
}
