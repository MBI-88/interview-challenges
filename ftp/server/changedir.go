package server

import (
	"net"
	"os"
	"path/filepath"
	"strings"
	"regexp"
)

// Cd method. Change dirs
func (f *FTP) cd(conn net.Conn, data []string) {
	re,_ := regexp.Compile("../")
	if len(data) != 1 {
		f.respond(conn, status501)
		return
	}
	if re.MatchString(data[0]) {
		f.respond(conn,status550)
		return
	}
	switch data[0] {
	case "..":
		if f.workdir == f.basedir {
			f.respond(conn, status550)
		} else {
			dirs := strings.Split(f.workdir, "/")
			dirs = dirs[:len(dirs)-1]
			f.workdir = strings.Join(dirs, "/")
			f.respond(conn, status250)
		}

	default:
		for ps, str := range strings.Split(f.workdir, "/") {
			if str == data[0] {
				tempdir := strings.Join(strings.Split(f.workdir, "/")[:ps+1], "/")
				if strings.Contains(f.basedir, tempdir) {
					f.workdir = tempdir
					f.respond(conn, status250)
				} else {
					f.respond(conn, status550)
				}
				return
			}
		}
		if _, err := os.Stat(filepath.Join(f.workdir, data[0])); os.IsNotExist(err) {
			f.respond(conn, status550)
			return
		}

		f.workdir = filepath.Join(f.workdir, data[0])
		f.respond(conn, status250)
	}

}
