package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
	"time"
)

const (
	status150 = "150 Opening data connection."
	status200 = "200 Command okay."
	status211 = "211 %s."
	status213 = "213 %d"
	status214 = "214 %s"
	status215 = "215 %s"
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status227 = "227 Entering Passive Mode (%s,%d,%d)"
	status230 = "230 User %s logged."
	status250 = "250 Directory successfully changed."
	status257 = "257 %q is the current directory."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status450 = "450 Requested file action not taken. File unavailable."
	status500 = "500 Bad port."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Cammand not implemented for that parameter."
	status522 = "522 EPSV not allowed. Please use PASV for data connections."
	status550 = "550 File or Dir not found or permission denied."
)

var commands = []string{
	"QUIT", "LIST", "CWD", "MKD", "RMD", "RETR", "STOR", "HELP", "PWD", "MGET",
	"MPUT", "DELE", "MDELE", "PASV", "USER", "SYST", "FEAT", "PORT", "EPSV", "SIZE", "TYPE", "NLST", "ABORT", "SITE"}

/*
	Commands: mdelete,mget don't work for the client ftp (server recibes nlst insted)
	mput works suing put as tool
*/

// FTP server
type FTP struct {
	basedir, workdir, remotepath, eofile, ip string
	por                                      int
}

// SetDir method to set up
func (f *FTP) SetDir(dir, ip string, por int) {
	basedir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	f.basedir = basedir
	f.workdir = basedir
	f.por = por
	f.ip = ip
}

// HandlerConnectControl method. Control connection port
func (f *FTP) handlerConnectControl(conn net.Conn) {
	signal := make(chan struct{})
	timer := 5 * time.Minute
	s := bufio.NewScanner(conn)
	f.eofile = "\r\n"
	f.respond(conn, status220)
	go func() {
		for s.Scan() {
			signal <- struct{}{}
			input := strings.Fields(s.Text())
			com, data := input[0], input[1:]
			log.Printf("<< %s %s", com, data)
			switch com {
			case commands[0]: // quit
				f.exit(conn)
			case commands[1]: // list
				f.ls(conn, data)
			case commands[2]: // cd
				f.cd(conn, data)
			case commands[3]: // mkdir
				f.mkdir(conn, data)
			case commands[4]: // rmdir
				f.rmdir(conn, data)
			case commands[5]: // get
				f.get(conn, data)
			case commands[6]: // put
				f.put(conn, data)
			case commands[7]: // help
				f.help(conn)
			case commands[8]: // pwd
				f.pwd(conn)
			case commands[9]: // mget
				f.mget(conn, data)
			case commands[10]: // mput
				f.mput(conn, data[1:])
			case commands[11]: // delete
				f.delete(conn, data)
			case commands[12]: // mdelete
				f.mdelete(conn, data)
			case commands[13]: // pasv
				f.pasv(conn)
			case commands[14]: // user
				f.user(conn, data)
			case commands[15]: // syst
				f.syst(conn)
			case commands[16]: // feat
				f.feat(conn)
			case commands[17]: // port
				f.port(conn, data)
			case commands[18]: // epsv
				f.epsv(conn)
			case commands[19]: // size
				f.size(conn, data)
			case commands[20]: // type
				f.conType(conn, data)
			case commands[21]: // nlst
				f.nlst(conn, data)
			case commands[22]: // abort
				f.exit(conn)
			case commands[23]: //site
				f.chmod(conn, data)
			default:
				f.respond(conn, status502)
			}

		}

	}()
	for {
		select {
		case _, ok := <-signal:
			if !ok {
				conn.Close()
				return
			}
		case <-time.After(timer):
			conn.Close()
			return
		}
	}

}

// StartServer method
func (f *FTP) StartServer() {
	fmt.Printf("[*] Running server on %s:%d\n[*] Serving file on %s\n", f.ip, f.por, f.basedir)
	endpoint := fmt.Sprintf("%s:%d", f.ip, f.por)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("[-] Error in setting connection %s\n", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Printf("[-] Error in connection %s\n", err)
		}
		fmt.Printf("[+] Accept connection IP:%s\n", conn.RemoteAddr())
		go f.handlerConnectControl(conn)

	}
}

func (f FTP) respond(conn net.Conn, s string) {
	s = s + f.eofile
	fmt.Fprintf(conn, s)

}

func (f FTP) connection() (net.Conn, error) {
	newconn, err := net.Dial("tcp", f.remotepath)
	return newconn, err
}
