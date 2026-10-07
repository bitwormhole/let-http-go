package lethttpgo

import (
	"net"

	"github.com/starter-go/afs"
)

type Context struct {
	Configuration *Configuration

	Addr *net.TCPAddr

	Root afs.Path
}
