package main

import "fmt"

func Print(message string) {
	fmt.Printf(message)
}

func GeneratePadding(contents string) string {
	if len(contents) == 0 {
		return " "
	}

	return ""
}

func PrintBoard(board [9]string) {
	Print(
		`
|==============|
|  ` + board[0] + GeneratePadding(board[0]) + ` |  ` + board[1] + GeneratePadding(board[1]) + ` |  ` + board[2] + GeneratePadding(board[2]) + ` |
|==============|
|  ` + board[3] + GeneratePadding(board[3]) + ` |  ` + board[4] + GeneratePadding(board[4]) + ` |  ` + board[5] + GeneratePadding(board[5]) + ` |
|==============|
|  ` + board[6] + GeneratePadding(board[6]) + ` |  ` + board[7] + GeneratePadding(board[7]) + ` |  ` + board[8] + GeneratePadding(board[8]) + ` |
|==============|

`)

}

func WipeConsole() {
	Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}
