package server

import (
	"net"
)

// Delete method. Delete a file
func (f FTP) delete(conn net.Conn, file []string) {
	if len(file) != 1 || file[0] == "" {
		f.respond(conn, status501)
		return

	}
	f.helperDel(conn,file[0])
}
