package server

import (
	"net"
)

// Get method. Download a file
func (f FTP) get(conn net.Conn, data []string) {
	if len(data) != 1 || data[0] == "" {
		f.respond(conn, status501)
		return
	}
	f.helperGet(conn,data[0])
}
