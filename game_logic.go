package main

func IsValidMove(move string, board [9]string) bool {
	if len(move) != 1 {
		return false
	}

	switch move {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		var index = GetBoardIndexForValidInput(move)

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

	// Horizontal rows.
	if board[0] != "" {
		if board[0] == board[1] && board[0] == board[2] {
			return true
		}
	}
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
