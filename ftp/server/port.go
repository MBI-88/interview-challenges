package server

import (
	"net"
)

// Port method. Set new remote connection
func (f *FTP) port(conn net.Conn, data []string) {
	remotehost := f.helperRemoteAddress(data)
	if remotehost == "" {
		f.respond(conn,status500)
		return
	}
	f.remotepath = remotehost
	f.respond(conn,status200)
}