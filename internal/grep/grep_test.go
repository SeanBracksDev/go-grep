package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestIsDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	tests := []struct {
		path     string
		expected bool
	}{
		{dir, true},
		{file.Name(), false},
	}

	for _, test := range tests {
		result, err := IsDir(test.path)
		if err != nil {
			t.Errorf("IsDir(%s) returned error: %v", test.path, err)
		}
		if result != test.expected {
			t.Errorf("Expected %v; Got %v", test.expected, result)
		}
	}
}

func TestSearchFile(t *testing.T) {
	// ? This doesn't test the color output, not sure how to do that, or if it's even necessary
	content := "Hello, world!\nThis is a test file.\nHello again!"
	file, err := os.CreateTemp("", "testfile")
	tmpFile := file.Name()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	if _, err := file.WriteString(content); err != nil {
		t.Fatal(err)
	}
	file.Close()

	tests := []struct {
		searchString string
		lineNumbers  bool
		expected     string
	}{
		{"Hello", false, fmt.Sprintf("%s:Hello, world!\n%s:Hello again!\n", tmpFile, tmpFile)},
		{"Hello", true, fmt.Sprintf("%s:1:Hello, world!\n%s:3:Hello again!\n", tmpFile, tmpFile)},
		{"test", false, fmt.Sprintf("%s:This is a test file.\n", tmpFile)},
		{"NotFound", false, ""},
	}

	for _, test := range tests {
		t.Run(test.searchString, func(t *testing.T) {
			var output strings.Builder
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			fileReader, err := os.Open(file.Name())
			if err != nil {
				panic(err)
			}
			defer fileReader.Close()

			SearchFile(file.Name(), fileReader, test.searchString, test.lineNumbers)
			w.Close()

			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				output.WriteString(scanner.Text() + "\n")
			}

			os.Stdout = old

			if output.String() != test.expected {
				t.Errorf("Expected: '%v'; Got: '%v'", test.expected, output.String())
			}
		})
	}
}
