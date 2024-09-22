package server


import (
	"fmt"
	"strings"
	"net"
)

// Help method. Show help
func (f FTP) help(conn net.Conn) {
	infoConmands := strings.Join(commands," ")
	if _,err := conn.Write([]byte(infoConmands+" ")); err != nil {
		f.respond(conn,status426)
		return
	}
	f.respond(conn,fmt.Sprintf(status214,"Help OK") )
	
}