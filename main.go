package main

/* COMMENT: Makes the indentation a bit weird but this is the best way to write text as it is meant to be
   shown IMO. We add some extra line breaks as padding. In hindsight I could make this a constant somewhere and just print it. */
func ShowIntroduction() {
	Print(`

Welcome to Noughts and Crosses! The best game ever made.
Each player will take turns entering their chosen move.
You can make your move by entering the numbers 1-9. 1-3 is the first row, 4-6 is the second, and so on. 1 is the top-left and 3 is the top-right.

Press enter to get started.`)
}

func main() {
	/* COMMENT: Obviously you wouldn't normally put business logic in main() but we might as well since this
	   is a simple application. */
	InitialiseInput()
	WipeConsole()
	ShowIntroduction()
	GetInput("")

	/* COMMENT: Not sure what kinds of naming conventions are going on in the Go community but this is generally
	   how variable names are formatted in C and Go is largely based on C, so I went with that. */
	var whose_turn string
	var input string
	var board[9] string

	for turn := 0; turn < 9 && !HasSomeoneWon(board); turn++ {
		/* COMMENT: Probably should be a space before and after the %. I think VS Code auto-formatting messed it up. */
		if turn%2 == 0 {
			whose_turn = "Player 1"
		} else {
			whose_turn = "Player 2"
		}

		WipeConsole()
		PrintBoard(board)

		input = ""
		input = GetInput(whose_turn + ", make your move")

		for !IsValidMove(input, board) {
			input = GetInput("Sorry, that move is invalid. Try again")
		}

		if turn%2 == 0 {
			board[GetBoardIndexForValidInput(input)] = "X"
		} else {
			board[GetBoardIndexForValidInput(input)] = "O"
		}

		/* COMMENT: If someone has suddenly won then it must be the person who just made a move. */
		if HasSomeoneWon(board) {
			WipeConsole()
			PrintBoard(board)
			Print(whose_turn + " has won!\nThanks for playing!\n")
		}
	}

	if !HasSomeoneWon(board) {
		WipeConsole()
		PrintBoard(board)
		Print("Nobody won 😢. Better luck next time!\n")
	}
}
