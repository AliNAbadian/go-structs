package main

import (
	"fmt"
	"strings"
)

func InputReader(promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')

	return strings.TrimSpace(userInput)
}
