package connect_four_game

import "testing"

func TestWonVerticalPlayer1(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	if wonVertical(board) != 1 {
		t.Error("win not detected")
	}
}

func TestWonVerticalPlayer2(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, -1, -1, -1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	if wonVertical(board) != -1 {
		t.Error("win not detected")
	}

}

func TestNoWinVertical(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, -1, -1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	if wonVertical(board) != 0 {
		t.Error("win not detected")
	}
}

func TestInterleavedVertical(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{1, 1, -1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	if wonVertical(board) != 0 {
		t.Error("win not detected")
	}
}

func TestPlayer1WinHorizontal(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	if wonHorizontal(board) != 1 {
		t.Fail()
	}
	board2 := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
	}
	if wonHorizontal(board2) != 0 {
		t.Fail()
	}


}

func TestWonTopRow(t *testing.T) {
	board3 := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, 1, -1, 1, -1, 1},
		{1, -1, -1, 1, -1, 1},
		{-1, 1, 1, -1, -1, 1},
		{1, -1, 1, 1, 1, 1},
	}
	if wonHorizontal(board3) != 1 {
		t.Error("Failed on top row")
	}
	board4 := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, 1, -1, 1, -1, 1},
		{1, -1, -1, 1, -1, 1},
		{-1, 1, 1, -1, -1, -1},
		{1, -1, 1, 1, 1, 1},
	}
	if wonHorizontal(board4) != 0 {
		t.Fail()
	}
}

func TestForwardDiagonalWin(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, 1, -1, 1, -1, 1},
		{1, -1, -1, 1, -1, 1},
		{-1, 1, -1, -1, -1, -1},
		{1, -1, 1, -1, -1, 1},
	}

	if res := wonDiagonalForward(board); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}

	board2 := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, 1, 1, 1, -1, 1},
		{1, -1, -1, 1, -1, 1},
		{-1, 1, -1, 1, 1, -1},
		{1, -1, 1, 1, -1, 1},
	}
	if res := wonDiagonalForward(board2); res != 1 {
		t.Errorf("Expected 1, got %d", res)
	}
}

func TestBackwardsDiagonalWin(t *testing.T) {
	board := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{-1, 1, -1, 1, -1, 1},
		{1, -1, 1, 1, -1, 1},
		{-1, 1, -1, -1, -1, -1},
		{1, -1, 1, -1, -1, 1},
	}

	if res := wonDiagonalBackwards(board); res != 1 {
		t.Errorf("Expected 1, got %d", res)
	}

	board2 := [][]int32{
		{},
		{},
		{},
		{-1, 1, -1, 1, -1, -1},
		{1, -1, 1, 1, -1, 1},
		{-1, 1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, 1},
	}
	if res := wonDiagonalBackwards(board2); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}
}

func TestJagged(t *testing.T) {
	board := [][]int32{
		{1, -1},
		{},
		{1, -1, 1},
		{1, 1, 1, -1},
		{},
		{1, -1, -1, -1},
		{},
	}

	allMethods := []func(board [][]int32)int32 {
		wonHorizontal,
		wonVertical,
		wonDiagonalBackwards,
		wonDiagonalForward,
	}

	for f := range allMethods {
		if res := allMethods[f](board); res != 0 {
			t.Fail()
		}
	}
}

func TestGetSerial(t *testing.T) {
	game := NewConnectFourGame()
	_ = game.Move(2, 1)
	_ = game.Move(1, -1)
	_ = game.Move(1, 1)
	_ = game.Move(5, -1)
	expected := []int32{0, 2, -1, 1, 1, 1, 0, 0, 1, -1, 0}
	if serial := game.GetSerial();  ! equal(serial, expected) {
		t.Errorf("Incorrect comparison: expected %d, got %d", expected, serial)
	}

}

func equal(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}