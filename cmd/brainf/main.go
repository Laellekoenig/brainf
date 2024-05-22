package main

import (
	"fmt"
	"os"

	"github.com/Laellekoenig/brainf/internal/brainfuck"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: brainf filename")
		return
	}

	fileName := os.Args[1]
	code, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to read file %s: %s\n", fileName, err)
		return
	}

	program := brainfuck.NewProgram(code)
	fmt.Print(string(program.Run()))
}
