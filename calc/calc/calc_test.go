package calc

import (
	"math/big"
	"testing"
)

func floatEqual(first, second float64) bool {
	return big.NewFloat(first).Cmp(big.NewFloat(second)) == 0
}

func TestCalculate(t *testing.T) {
	t.Run("Same priority", func(t *testing.T) {
		expr := "1 + 2 - 3"
		expected := 0.0

		got, _ := Calculate(expr)

		if !floatEqual(expected, got) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Different priorities", func(t *testing.T) {
		expr := "2 + 2 * 2"
		expected := 6.0

		got, _ := Calculate(expr)

		if !floatEqual(expected, got) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("With brackets", func(t *testing.T) {
		expr := "(1 + 2) * 3"
		expected := 9.0

		got, _ := Calculate(expr)

		if !floatEqual(expected, got) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})

	t.Run("Mixed float operations", func(t *testing.T) {
		expr := "2 * ((3 + 5) * 10 - 35 / 23) * 33 + 12"
		expected := 5191.565217391304

		got, _ := Calculate(expr)

		if !floatEqual(expected, got) {
			t.Error("\nExpected: ", expected, "\nGot: ", got)
		}
	})
}
