package lib

import (
	"bufio"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	// The data file is put into the directory relative to the *runfiles
	// dir, in a directory corresponding to the package name.
	file, err := os.Open("file.txt")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if text := scanner.Text(); text != SayHello() {
			t.Errorf("Mismatch: '%v'", text)
		}
	}

}
