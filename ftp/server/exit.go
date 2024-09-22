package server


import (
	"net"
)

// Exit method. Close connection
func (f FTP) exit(conn net.Conn) {
	f.respond(conn,status221)
 	defer conn.Close()
}