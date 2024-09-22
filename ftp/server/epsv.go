package server

import (
	"net"
)

// Epsv method. Return a ftp listening port
func (f *FTP) epsv(conn net.Conn){
	f.respond(conn,status522)
}