package server

import (
	"net"
)

// Mget method. Download files
func (f FTP) mget(conn net.Conn, datas []string) {
	for _, data := range datas {
		if data == "" {
			f.respond(conn,status501)
			continue
		}
		f.helperGet(conn,data)
	}

}
