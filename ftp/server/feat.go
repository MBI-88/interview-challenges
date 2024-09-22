package server

import (
	"fmt"
	"net"
)

// Feat method. Return all characters
func (f FTP) feat(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf(status211,"-Features")))
	for _, item := range []string{"UTF8","SIZE"}{
		if _, err := conn.Write([]byte(item)); err != nil {
			f.respond(conn, status426)
			return
		}
	}
	f.respond(conn, fmt.Sprintf(status211, " End"))

}
