package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var InputReader *bufio.Reader

/* COMMENT: No point initialising this over and over again, so we do it once at program startup. */
func InitialiseInput() {
	InputReader = bufio.NewReader(os.Stdin)
}

/* COMMENT: Wrapper to make my life easier and make the program easier to read. */
func GetInput(prefix_message string) string {
	if len(prefix_message) > 0 {
		fmt.Print(prefix_message + ": ")
	}

	/* COMMENT: I don't really care about checking for errors etc since this is a simple application. */
	text, _ := InputReader.ReadString('\n')

	/* COMMENT: These could have been combined into one line. */
	// Remove last char, which is a newline.
	text = text[:len(text)-1]
	return text
}

/* COMMENT: If you wanted to be really safe you could double-check the input is valid here. */
func GetBoardIndexForValidInput(input string) int {
	var result, _ = strconv.Atoi(input)
	return result - 1
}
