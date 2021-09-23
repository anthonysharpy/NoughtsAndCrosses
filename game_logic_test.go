package main

import (
	"strings"
	"testing"
)

/* COMMENT: The brief for this said that the unit tests should expect successes and failures, but I'm
   not sure what was meant by that, since in theory unit tests should always be successes (if they all pass).
   I'm assuming it meant to test both possibilities, with I did with the victory/no victory tests. I could have
   written lots more tests here, but not sure what that would have really demonstrated. */

func TestIsValidMove_TryGibberish(t *testing.T) {
	var board [9]string

	/* COMMENT: In hindsight it might have been better to declare an array of "moves" and then go through them all with a for loop. */
	var move string = "lol"
	if IsValidMove(move, board) {
		/* COMMENT: Missing space at beginning of string here. Not that it really matters since it's a unit test. */
		t.Error(move + "should not have been a valid move.")
	}

	move = "0"
	if IsValidMove(move, board) {
		t.Error(move + "should not have been a valid move.")
	}

	move = "10"
	if IsValidMove(move, board) {
		t.Error(move + "should not have been a valid move.")
	}

	move = "444"
	if IsValidMove(move, board) {
		t.Error(move + "should not have been a valid move.")
	}

	move = "sdfw45bw.5b24305bh34"
	if IsValidMove(move, board) {
		t.Error(move + "should not have been a valid move.")
	}
}

func TestHasSomeoneWon_Victories(t *testing.T) {

	/* COMMENT: In hindsight it might have been better to declare an array of boards and then go through them all with a for loop. */
	
	// Test horizontal victories.
	var board = [9]string{"X", "X", "X", "", "", "", "", "", ""}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "", "", "O", "O", "O", "", "", ""}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "", "", "", "", "", "X", "X", "X"}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	// Test vertical victories.
	board = [9]string{"X", "", "", "X", "", "", "X", "", ""}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "O", "", "", "O", "", "", "O", ""}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "", "O", "", "", "O", "", "", "O"}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	// Test diagnonal victories.
	board = [9]string{"X", "", "", "", "X", "", "", "", "X"}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "", "O", "", "O", "", "O", "", ""}
	if !HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned true with: " + strings.Join(board[:], ", "))
	}
}

/* Test random cases where there should be no victory. */
func TestHasSomeoneWon_NoVictory(t *testing.T) {

	var board = [9]string{"X", "X", "O", "O", "X", "X", "X", "O", "O"}
	if HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned false with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"", "", "", "", "", "", "", "", ""}
	if HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned false with: " + strings.Join(board[:], ", "))
	}

	board = [9]string{"O", "O", "X", "X", "O", "O", "O", "X", "X"}
	if HasSomeoneWon(board) {
		t.Error("HasSomeoneWon should have returned false with: " + strings.Join(board[:], ", "))
	}
}
