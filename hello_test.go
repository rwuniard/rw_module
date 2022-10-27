package rw_module

import "testing"

func TestSayHello(t *testing.T) {
	result := SayHello()
	expected_result := "Hello this is from rw_module."

	if result != expected_result {
		t.Error("Supposed to get:", expected_result, "but got:", result)
	}
}
