package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"enlace/hamming"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sErro ao ler a entrada: %v%s\n", Red, err, Reset)
		os.Exit(1)
	}

	frame := strings.TrimSpace(input)

	header := "01111110"
	terminator := "01111110"

	if !strings.HasPrefix(frame, header) || !strings.HasSuffix(frame, terminator) {
		fmt.Fprintf(os.Stderr, "%sFrame inválido%s\n", Red, Reset)
		os.Exit(1)
	}

	content := frame[len(header) : len(frame)-len(terminator)]

	runes := []rune(content)
	if len(runes) > 3 {
		if runes[3] == '0' {
			runes[3] = '1'
		} else {
			runes[3] = '0'
		}
	}
	contentWithError := string(runes)
	fmt.Printf("%sContent with error: %s%s\n", Blue, contentWithError, Reset)

	correctedCode := hamming.CorrectHammingCode(contentWithError)
	fmt.Printf("%sCorrected code: %s%s\n", Green, correctedCode, Reset)

	data := hamming.ExtractData(correctedCode)
	fmt.Printf("%sExtracted data: %s%s\n", Yellow, data, Reset)
}
