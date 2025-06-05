package main

import "testing"

// func TestReturnGreaterNum(t *testing.T) {
// 	result, err := ReturnGreaterNum(1, 2)
// 	if err != nil {
// 		t.Errorf("Error in test: %v", err)
// 	}
// 	if result != 2 {
// 		t.Errorf("Result was incorrect, got %d, want %d", result, 2)
// 	}
// }

func TestReturnGreaterNum(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped in short mode")
	}
	var tests = []struct {
		name  string
		input struct {
			a int
			b int
		}
		want int
	}{
		{"9 should be greater that 0", struct {
			a int
			b int
		}{9, 3}, 9},
		{"5 should be greater that 2", struct {
			a int
			b int
		}{5, 2}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // run tests in parallel
			result, err := ReturnGreaterNum(tt.input.a, tt.input.b)
			if err != nil {
				t.Errorf("Error in test: %v", err)
			}
			if result != tt.want {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}
