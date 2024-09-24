package grep

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/SeanBracksDev/go-grep/internal/colour"
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

func Search(input io.Reader, searchString []byte, opts ...Option) {
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
		line := scanner.Bytes()
		if bytes.Contains(line, searchString) {
			var sb strings.Builder

			if options.filePath != nil && *options.filePath != "" {
				sb.Write(colour.Colour(*options.filePath, colour.Blue))
				sb.WriteByte(58)
			}
			if options.lineNumbers != nil && *options.lineNumbers {
				sb.Write(colour.Colour(lineNumber, colour.Green))
				sb.WriteByte(58)
			}
			sb.Write(bytes.ReplaceAll(line, searchString, colour.Colour(searchString, colour.Red))) // ? this is pretty expensive, is there a better way?
			fmt.Println(sb.String())
		}
		lineNumber++
	}
}
