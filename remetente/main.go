package main

import (
	"fmt"
	"os"

	"enlace/hamming"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <sequencia de bits>\n", os.Args[0])
		os.Exit(1)
	}

	data := os.Args[1]
	if !hamming.IsValidBinary(data) {
		fmt.Fprintf(os.Stderr, "A sequência de bits deve conter apenas 0s e 1s\n")
		os.Exit(1)
	}

	header := "01111110"
	terminator := "01111110"

	hammingCode := hamming.GenerateHammingCode(data)

	frame := header + hammingCode + terminator

	fmt.Println(frame)
}
