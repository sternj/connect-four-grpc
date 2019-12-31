package connect_four_server

import (
	"context"
	"github.com/google/uuid"
	connect_four_game "github.com/sternj/go-connect-four/internal/connect-four-game"
	connect_four_proto "github.com/sternj/go-connect-four/internal/connect-four-proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"fmt"
)

func (c *ConnectFourService) StartGame(ctx context.Context,
	in *connect_four_proto.StartGameRequest) (
	*connect_four_proto.StartGameResponse,
	error) {

		fmt.Println("START GAME REQ")
	gameUuid := uuid.New().String()

	c.gameLocks[gameUuid] = &sync.Mutex{}
	c.games[gameUuid] = connect_four_game.NewConnectFourGame()

	return &connect_four_proto.StartGameResponse{
		Started:              true,
		Uuid:                 gameUuid,
	}, nil
}


func (c *ConnectFourService) Move(ctx context.Context,
	in *connect_four_proto.MoveRequest) (
	*connect_four_proto.MoveResponse,
	error) {
		gameUuid := in.GetGameUuid()
		lock, ok := c.gameLocks[gameUuid]
		if ! ok {
			return nil, status.Error(codes.NotFound, "game not found")
		}
		lock.Lock()
		defer c.gameLocks[gameUuid].Unlock()
		game := c.games[gameUuid]
		if err := game.Move(in.GetColumn(), in.GetPlayer()); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if winner := game.Won(); winner != 0 {
			return &connect_four_proto.MoveResponse{Board: game.GetSerial(), Winner: winner}, nil
		}

		return &connect_four_proto.MoveResponse{Board: game.GetSerial(), Winner: 0}, nil
}