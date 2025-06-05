package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFileOperations(t *testing.T) {
	filename := "testfile"
	file, err := os.Create(filename)
	if err != nil {
		t.Errorf("failed to create file: %v", err)
	}

	t.Cleanup(func() {
		err = os.Remove(filename)
		if err != nil {
			fmt.Printf("Error removing file: %v\n", err)
		}
	})

	_, err = file.WriteString("random string")
	if err != nil {
		t.Errorf("failed to write string to file: %v", err)
	}
}
