package tdd

import "testing"

func TestReverseString(t *testing.T) {
	input := "abc"
	expected := "cba"
	result := ReverseString(input) // This line will cause a compilation error
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
