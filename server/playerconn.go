package server

import (
	"gomud/entity"
	"net"
)

type PlayerConn struct {
	Player *entity.Player
	Conn   net.Conn
}
