package connect_four_game

import "fmt"

type ConnectFour struct {
	board [][]int32
}

func NewConnectFourGame() *ConnectFour{
	return &ConnectFour{board: [][]int32{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}}
}

func (c *ConnectFour) Move(column int32, player int32) error{
	if ValidMove(c.board, column, player) {
		return nil
	}
	return fmt.Errorf("invalid move in col %d", column)
}

func (c *ConnectFour) GetSerial() []int32 {
	serialBoard := make([]int32, 0)
	for col := range c.board {
		serialBoard = append(serialBoard, int32(len(c.board[col])))
		for elem := range c.board[col] {
			serialBoard = append(serialBoard, c.board[col][elem])
		}
	}
	return serialBoard
}

func (c *ConnectFour) Won() int32 {
	winFuncs := []func([][]int32)int32 {
		wonHorizontal,
		wonVertical,
		wonDiagonalForward,
		wonDiagonalBackwards,
	}
	for f := range winFuncs {
		if won := winFuncs[f](c.board); won != 0 {
			return won
		}
	}

	return 0
}