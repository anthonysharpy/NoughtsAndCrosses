package main

func ShowIntroduction() {
	Print(`

Welcome to Noughts and Crosses! The best game ever made.
Each player will take turns entering their chosen move.
You can make your move by entering the numbers 1-9. 1-3 is the first row, 4-6 is the second, and so on. 1 is the top-left and 3 is the top-right.

Press enter to get started.`)
}

var Board [9]string

func main() {
	InitialiseInput()
	WipeConsole()
	ShowIntroduction()
	GetInput("")

	var whoseTurn string
	var input string

	for turn := 0; turn < 9 && !HasSomeoneWon(Board); turn++ {
		if turn%2 == 0 {
			whoseTurn = "Player 1"
		} else {
			whoseTurn = "Player 2"
		}

		WipeConsole()
		PrintBoard(Board)

		input = ""
		input = GetInput(whoseTurn + ", make your move")

		for !IsValidMove(input, Board) {
			input = GetInput("Sorry, that move is invalid. Try again")
		}

		if turn%2 == 0 {
			Board[GetBoardIndexForValidInput(input)] = "X"
		} else {
			Board[GetBoardIndexForValidInput(input)] = "O"
		}

		if HasSomeoneWon(Board) {
			WipeConsole()
			PrintBoard(Board)
			Print(whoseTurn + " has won!\nThanks for playing!\n")
		}
	}

	if !HasSomeoneWon(Board) {
		Print("Nobody won 😢. Better luck next time!\n")
	}
}
