package files

import (
	"bufio"
	"fmt"
	"io"
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

func SearchFile(filePath string, input io.Reader, searchString string, lineNumbers bool) {
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

func SearchStdin(input io.Reader, searchString string, lineNumbers bool) {
	scanner := bufio.NewScanner(input)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			var sb strings.Builder
			newLine := strings.ReplaceAll(line, searchString, highlightColour(searchString))
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
