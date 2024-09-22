package server

import (
	"net"
)

// Pasv method. Return not supported
func (f FTP) pasv(conn net.Conn){
	f.respond(conn,status502)
	
}