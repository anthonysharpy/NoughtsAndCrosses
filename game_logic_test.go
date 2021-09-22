package main

import (
	"strings"
	"testing"
)

func TestIsValidMove_TryGibberish(t *testing.T) {
	var board [9]string

	var move string = "lol"
	if IsValidMove(move, board) {
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
