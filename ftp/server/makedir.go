package server

import (
	"net"
)

// Mkdir method. Make a dir
func (f FTP) mkdir(conn net.Conn, dirs []string) {
	if len(dirs) != 1 {
		f.respond(conn,status501)
		return
	}
	f.helperMkdir(conn,dirs[0])
}
