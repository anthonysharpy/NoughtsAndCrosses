package main

import "fmt"

// COMMENT: Simple wrapper for fmt.Printf so I don't have to type as much.
func Print(message string) {
	fmt.Printf(message)
}

func GeneratePadding(contents string) string {
	if len(contents) == 0 {
		return " "
	}

	return ""
}

/* COMMENT: This isn't pretty but I can't think of any substantially better ways to do it. */
func PrintBoard(board [9]string) {
	Print(`
|==============|
|  ` + board[0] + GeneratePadding(board[0]) + ` |  ` + board[1] + GeneratePadding(board[1]) + ` |  ` + board[2] + GeneratePadding(board[2]) + ` |
|==============|
|  ` + board[3] + GeneratePadding(board[3]) + ` |  ` + board[4] + GeneratePadding(board[4]) + ` |  ` + board[5] + GeneratePadding(board[5]) + ` |
|==============|
|  ` + board[6] + GeneratePadding(board[6]) + ` |  ` + board[7] + GeneratePadding(board[7]) + ` |  ` + board[8] + GeneratePadding(board[8]) + ` |
|==============|

`)

}

/* COMMENT: There is no cross-platform way to clear a terminal, so I didn't even bother trying and
   just went with this instead. It's possible to use something different depending on the detected OS
   to clear the terminal, but that just seemed overkill for a simple console application like this. */
func WipeConsole() {
	Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}
