package submodule

import "testing"

func TestMultiply(t *testing.T) {
	a := []int{1, 2, 3}

	result := Multiply(a...)
	expected_result := 6

	if result != expected_result {
		t.Error("Expected result:", expected_result, "but got:", result)
	}
}
