package server

import (
	"net"
)

// Put method. Upload a file
func (f FTP) put(conn net.Conn, data []string) {
	if len(data) != 1 || data[0] == "" {
		f.respond(conn,status501)
		return
	}
	f.helperPut(conn,data[0])
}
