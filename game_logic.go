package main

func IsValidMove(move string, board [9]string) bool {
	if len(move) != 1 {
		return false
	}

	switch move {
	/* COMMENT: Could have done some fancy tricks by converting the string to a byte and evalauting it
	   as an integer, but this way is much easier to read (and write). */
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		var index = GetBoardIndexForValidInput(move)

		/* COMMENT: This could have been greatly simplified by doing "return len(board[index]) == 0". Oops! */
		if len(board[index]) != 0 {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}


func HasSomeoneWon(board [9]string) bool {

	/* COMMENT: This could easily be shortened with some for loops, but I think that would make the code
	   unjustifiably harder to read. Doing it this way emphasises that there are exactly 8 ways to win a
	   game of noughts and crosses. */

	// Horizontal rows.

	if board[0] != "" {
		if board[0] == board[1] && board[0] == board[2] {
			return true
		}
	}
	/* COMMENT: Could have made these into one if statement, but I wanted to strongly emphasise that these are two
	   different layers of logic. */
	if board[3] != "" {
		if board[3] == board[4] && board[3] == board[5] {
			return true
		}
	}
	if board[6] != "" {
		if board[6] == board[7] && board[6] == board[8] {
			return true
		}
	}

	// Vertical rows.
	if board[0] != "" {
		if board[0] == board[3] && board[0] == board[6] {
			return true
		}
	}
	if board[1] != "" {
		if board[1] == board[4] && board[1] == board[7] {
			return true
		}
	}
	if board[2] != "" {
		if board[2] == board[5] && board[2] == board[8] {
			return true
		}
	}

	// Diagonals.
	if board[0] != "" {
		if board[0] == board[4] && board[0] == board[8] {
			return true
		}
	}
	if board[2] != "" {
		if board[2] == board[4] && board[2] == board[6] {
			return true
		}
	}

	return false
}
