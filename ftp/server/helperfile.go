package server

import (
	"path/filepath"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

func (f FTP) helperPut(conn net.Conn, filename string) {
	file, err := os.Create(filepath.Join(f.workdir,filename))
	if err != nil {
		f.respond(conn, status550)
		return
	}
	defer file.Close()
	wconn, err := f.connection()
	defer wconn.Close()
	if err != nil {
		f.respond(conn, status425)
		return
	}
	f.respond(conn, status150)
	buffer := make([]byte, 2048)
	for {
		n, err := wconn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				f.respond(conn, status226)
				return
			}
			f.respond(conn, status426)
			return
		}
		if _, err := file.Write(buffer[:n]); err != nil {
			f.respond(conn, status426)
			return
		}

	}
}

func (f FTP) helperGet(conn net.Conn, data string) {
	info, err := os.Stat(filepath.Join(f.workdir,data))
	if err != nil {
		f.respond(conn, status550)
		return
	}
	file, err := os.OpenFile(filepath.Join(f.workdir,data), os.O_RDONLY, 0775)
	defer file.Close()
	if err != nil {
		f.respond(conn, status550)
		return
	}

	filename := make([]byte, 1024)
	for _, el := range info.Name() {
		filename = append(filename, byte(el))
	}
	f.respond(conn, status150)

	wconn, er := f.connection()
	defer wconn.Close()
	if er != nil {
		f.respond(conn, status425)
		return
	}
	if _, err := wconn.Write(filename); err != nil {
		f.respond(conn, status426)
		return
	}

	if info.Size() >= 1e+9 {
		block := int64(1024 * 1024)
		buffer := make([]byte, block)
		for {
			b, err := file.Read(buffer)
			if err == io.EOF {
				wconn.Write([]byte(f.eofile))
				f.respond(conn, status226)
				return
			}
			if err != nil {
				f.respond(conn, status550)
				return
			}
			if b > 0 {
				if _, err := io.CopyN(wconn, bytes.NewReader(buffer[:b]), block); err != nil {
					f.respond(conn, status426)
					return
				}

			}

		}

	}
	_, e := io.Copy(wconn, file)
	if e != nil {
		f.respond(conn, status426)
		return
	}
	wconn.Write([]byte(f.eofile))
	f.respond(conn, status200)
}

func (f FTP) helperDel(conn net.Conn, file string) {
	if err := os.Remove(filepath.Join(f.workdir,file)); os.IsExist(err) {
		f.respond(conn, status550)
	}
	f.respond(conn, status226)
}

func (f FTP) helperRemoteAddress(data []string) string {
	address := data[0]
	host := strings.Split(address, ",")
	ip := strings.Join(host[:4], ".")
	porthigh, err := strconv.Atoi(host[4])
	if err != nil {
		return ""
	}
	portlow, err := strconv.Atoi(host[len(host)-1])
	if err != nil {
		return ""
	}
	port := porthigh<<8 + portlow
	remotehost := fmt.Sprintf("%s:%d", ip, port)
	if err != nil {
		return ""
	}
	return remotehost

}

func (f FTP) helperList(conn net.Conn, path string) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		f.respond(conn, status550)
		return
	}
	f.respond(conn, status150)
	wconn, err := f.connection()
	if err != nil {
		f.respond(conn, status425)
		return
	}
	defer wconn.Close()
	for _, dir := range dirs {
		if _, err := fmt.Fprintf(wconn, "%s  %s %s %d bytes%s", dir.Mode(), dir.ModTime(), dir.Name(), dir.Size(), f.eofile); err != nil {
			f.respond(conn, status426)
			return
		}
	}
	f.respond(conn, status226)

}

func (f FTP) helperNlst(conn net.Conn, path string) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		f.respond(conn, status550)
		return
	}
	for _, item := range dirs {
		if _, er := fmt.Fprintf(conn, "%s ", item.Name()); er != nil {
			f.respond(conn, status426)
		}
	}
	f.respond(conn, status200)
}

func (f FTP) helperMkdir(conn net.Conn, data string) {
	if _, err := os.Stat(filepath.Join(f.workdir,data)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(f.workdir,data), 0755); os.IsExist(err) {
			f.respond(conn, status550)
			return
		}
	}

	f.respond(conn, status200)
}

func (f FTP) helperRmdir(conn net.Conn, data string) {
	if err := os.RemoveAll(filepath.Join(f.workdir,data)); os.IsExist(err) {
		f.respond(conn, status550)
		return
	}
	f.respond(conn, status200)
}

func (f FTP) helperChmod(conn net.Conn, perm string, data []string) {
	osperm, err := strconv.Atoi(perm)
	if err != nil {
		f.respond(conn, status550)
		return
	}
	for _, item := range data {
		if err := os.Chmod(filepath.Join(f.workdir,item), os.FileMode(osperm)); os.IsExist(err) {
			f.respond(conn, status550)
		}
	}
	f.respond(conn, status200)
}