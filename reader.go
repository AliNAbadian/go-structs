package main

import (
	"fmt"
	"strings"
)

func inputReader(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Trim(userInput, "\n")

	return cleanedInput
}
