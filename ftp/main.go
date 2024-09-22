package main

import (
	"flag"
	"fmt"
	"ftp/server"
	"os"
)

var (
	port *int
	help *string
	dir *string
	ip *string
)

func init() {
	port = flag.Int("P", 21, "port")
	help = flag.String("h","","show information")
	dir = flag.String("D","./","set shared dir")
	ip = flag.String("H","localhost","host")
	flag.Usage = func() {
		info := "[+] Commands:\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	ftp := server.FTP{}
	ftp.SetDir(*dir,*ip,*port)
	ftp.StartServer()
	

}


