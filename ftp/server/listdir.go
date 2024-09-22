package server

import (
	"net"
	"os"
	"path/filepath"
	"regexp"
)

// Ls method. List dirs
func (f FTP) ls(conn net.Conn, arg []string) {
	re, _ := regexp.Compile("../")
	if len(arg) != 0 && !re.MatchString(arg[0]) {
		if _, err := os.Stat(filepath.Join(f.workdir, arg[0])); os.IsNotExist(err) {
			f.respond(conn, status550)
			return
		}
		tempPaht := filepath.Join(f.workdir, arg[0])
		f.helperList(conn, tempPaht)
	} else {
		f.helperList(conn, f.workdir)

	}

}
