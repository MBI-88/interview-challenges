package server

import (
	"net"
)

// Mput method. Upload files
func (f FTP) mput(conn net.Conn, data []string) {
	for _, file := range data {
		if file == "" {
			f.respond(conn,status501)
			continue
		}
		f.helperPut(conn,file)
	}
}

