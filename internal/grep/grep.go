package grep

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

type Options struct {
	lineNumbers *bool
	filePath    *string
}

type Option func(options *Options) error

func WithLineNumbers() Option {
	return func(options *Options) error {
		withLineNumbers := true
		options.lineNumbers = &withLineNumbers
		return nil
	}
}

func WithFilePath(filePath string) Option {
	return func(options *Options) error {
		options.filePath = &filePath
		return nil
	}
}

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

func Search(input io.Reader, searchString string, opts ...Option) {
	var options Options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			panic(err)
		}
	}

	scanner := bufio.NewScanner(input)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			var sb strings.Builder
			newLine := strings.ReplaceAll(line, searchString, highlightColour(searchString))

			if options.filePath != nil && *options.filePath != "" {
				sb.WriteString(fmt.Sprintf("%s:", fileNameColour(*options.filePath)))
			}
			if options.lineNumbers != nil && *options.lineNumbers {
				sb.WriteString(fmt.Sprintf("%s:", lineNumberColour(lineNumber)))
			}
			sb.WriteString(newLine)
			fmt.Println(sb.String())
		}
		lineNumber++
	}
}
