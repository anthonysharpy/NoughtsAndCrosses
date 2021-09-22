package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var InputReader *bufio.Reader

func InitialiseInput() {
	InputReader = bufio.NewReader(os.Stdin)
}

func GetInput(prefix_message string) string {
	if len(prefix_message) > 0 {
		fmt.Print(prefix_message + ": ")
	}

	text, _ := InputReader.ReadString('\n')

	// Remove last char, which is a newline
	text = text[:len(text)-1]
	return text
}

func GetBoardIndexForValidInput(input string) int {
	var result, _ = strconv.Atoi(input)
	return result - 1
}
