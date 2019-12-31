package connect_four_server

import (
	connectfourgame "github.com/sternj/go-connect-four/internal/connect-four-game"
	connectfourproto "github.com/sternj/go-connect-four/internal/connect-four-proto"
	"sync"
)

type ConnectFourService struct {
	connectfourproto.UnimplementedConnectFourServer
	games map[string]*connectfourgame.ConnectFour
	gameLocks map[string]*sync.Mutex
}

func NewConnectFourService() *ConnectFourService {
	return &ConnectFourService{
		games:                          make(map[string]*connectfourgame.ConnectFour),
		gameLocks:                      make(map[string]*sync.Mutex),
	}
}