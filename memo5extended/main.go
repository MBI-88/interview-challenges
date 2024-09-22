package main

import (
	"fmt"
	"log"
	"time"
	"memo/memo5"
	"io/ioutil"
	"net/http"
)

func httpGetBody(url string,ch <-chan struct{}) (interface{}, error) {
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = ch
	resp, err := http.DefaultClient.Do(req)
	if err != nil {return nil,err}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	chanel := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			chanel <- url
		}
		close(chanel)
	}()
	return chanel
}



func main() {
	done := make(chan struct{})
	go func(){
		time.Sleep(2 * time.Second)
		done <-struct{}{}
	}()
	memo := memo5.New(httpGetBody)
	for url := range incomingURLs() {
		start := time.Now()
		value,err := memo.Get(url,done)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s %s %d\n",url,time.Since(start),len(value.([]byte)))
	}
}
