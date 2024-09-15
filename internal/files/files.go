package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	fileNameColour   = color.New(color.FgBlue).SprintFunc()
	highlightColour  = color.New(color.FgRed).SprintFunc()
	lineNumberColour = color.New(color.FgGreen).SprintFunc()
)

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fileInfo.IsDir() {
		return true, nil
	}
	return false, nil
}

// take a io.Reader and a string to search for
func Search(filePath, searchString string, lineNumbers bool) {
	input, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			var sb strings.Builder
			newLine := strings.ReplaceAll(line, searchString, highlightColour(searchString))
			sb.WriteString(fmt.Sprintf("%s:", fileNameColour(filePath)))
			if lineNumbers {
				sb.WriteString(fmt.Sprintf("%s:", lineNumberColour(lineNumber)))
			}
			sb.WriteString(newLine)
			fmt.Println(sb.String())
		}
		lineNumber++
	}
}

func BoyerMooreSearch(input, searchString string) {

}
