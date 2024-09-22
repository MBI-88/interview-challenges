package server

import (
	"net"
)

// Mdelete method. Delete files
func (f FTP) mdelete(conn net.Conn, files []string) {
	for _, file := range files {
		if file == "" {
			f.respond(conn,status501)
			continue
		}
		f.helperDel(conn,file)
	}
}
