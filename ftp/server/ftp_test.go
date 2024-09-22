package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"testing"
)

const (
	por = 4000
	dir = "./test"
	ip  = "localhost"
)

func clientAndserver() (net.Conn, net.Conn, FTP) {
	server := FTP{}
	server.SetDir(dir, ip, por)
	resp, req := net.Pipe()
	return resp, req, server
}

func TestVaribles(t *testing.T) {
	_, _, server := clientAndserver()
	if server.ip != ip {
		t.Errorf("[-] Error %s != %s\n", server.ip, ip)
	}
	if server.por != por {
		t.Errorf("[-] Error %d != %d\n", server.por, por)
	}
	testdir := strings.Split(server.workdir, "/")
	if fmt.Sprintf("./%s", testdir[len(testdir)-1]) != dir {
		t.Errorf("[-] Error %s != %s\n", server.workdir, dir)
	}
}

func TestCWD(t *testing.T) {
	response, req, server := clientAndserver()
	fmt.Println(server.workdir)
	go func() {
		defer req.Close()
		server.cd(req, []string{"../"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status550 {
		t.Errorf("[-] Error %s != status550\n", resp)
	}
	response.Close()
}

func TestCWD2(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.cd(req, []string{".."})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status550 {
		t.Errorf("[-] Error %s != status550\n", resp)
	}
	response.Close()
}

func TestCWD3(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.cd(req, []string{"../../"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status550 {
		t.Errorf("[-] Error %s != status550\n", resp)
	}
	response.Close()
}

func TestCWD4(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.cd(req, []string{"test1","test2"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status501 {
		t.Errorf("[-] Error %s != status501\n", resp)
	}
	response.Close()
}

func TestDelete(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.delete(req, []string{"Cuentas_Mensuales.txt"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status226 {
		t.Errorf("[-] Error %s != status226\n", resp)
	}
	response.Close()
}

func TestDelete2(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.delete(req, []string{""})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status501 {
		t.Errorf("[-] Error %s != status501\n", resp)
	}
	response.Close()
}

func TestHelp(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.help(req)

	}()
	status, _ := ioutil.ReadAll(response)
	if !strings.Contains(string(status), fmt.Sprintf(status214, "Help OK")) {
		t.Errorf("[-] Error %s != status214\n", status)
	}
	response.Close()
}

func TestMKD(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.mkdir(req, []string{"mydir"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status200 {
		t.Errorf("[-] Error %s != status200\n", resp)
	}
	response.Close()
}

func TestRMD(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.rmdir(req, []string{"mydir"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != status200 {
		t.Errorf("[-] Error %s != status200\n", resp)
	}
	response.Close()
}

func TestType(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.conType(req, []string{"A"})
	}()
	resp, _ := ioutil.ReadAll(response)
	status := strings.Fields(string(resp))[0]
	if status != "200" {
		t.Errorf("[-] Error %s != 200\n", resp)
	}
	if server.eofile != "\r\n" {
		t.Errorf("[-] Error eofile != A")
	}
	response.Close()
}

func TestType2(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.conType(req, []string{"I"})
	}()
	resp, _ := ioutil.ReadAll(response)
	status := strings.Fields(string(resp))[0]
	if status != "200" {
		t.Errorf("[-] Error %s != 200\n", resp)
	}
	if server.eofile != "\n" {
		t.Errorf("[-] Error eofile != I")
	}
	response.Close()
}

func TestUser(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.user(req, []string{"John"})
	}()
	resp, _ := ioutil.ReadAll(response)
	if string(resp) != fmt.Sprintf(status230, "John") {
		t.Errorf("[-] Error %s != status230\n", resp)
	}
	response.Close()
}

func TestSyst(t *testing.T) {
	response, req, server := clientAndserver()
	go func() {
		defer req.Close()
		server.syst(req)
	}()
	resp, _ := ioutil.ReadAll(response)
	status := strings.Fields(string(resp))[0]
	if status != "215" {
		t.Errorf("[-] Error %s != status215\n", status)
	}
	response.Close()
}

func TestPWD(t *testing.T) {
	response, req, server := clientAndserver()
	go func(){
		defer req.Close()
		server.pwd(req)
	}()
	resp,_ := ioutil.ReadAll(response)
	status := strings.Fields(string(resp))[0]
	if status != "257" {
		t.Errorf("[-] Error %s != status257\n",resp)
	}
}