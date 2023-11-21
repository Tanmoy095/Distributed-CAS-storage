package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	ListenAddress string
	Listener      net.Addr
	mu            sync.RWMutex //this mutex will protect peers
	peers         map[net.Addr]Peer
}

func NewTCPTransport(ListenAddress string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: ListenAddress,
	}
}
