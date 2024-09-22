package server

import (
	"regexp"
	"path/filepath"
	"net"
)


func (f FTP) nlst(conn net.Conn, data []string) {
	re,_ := regexp.Compile("../")
	if len(data) != 0 && !re.MatchString(data[0]) {
		path := filepath.Join(f.workdir,data[0])
		f.helperNlst(conn,path)
	}
	f.helperNlst(conn,f.workdir)
}