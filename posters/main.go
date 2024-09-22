package main

import (
	"flag"
	"fmt"
	"os"
	"omd/poster"
)

var (
	d  *bool
	args string
)

func init() {
	d = flag.Bool("d", false, "Download movie poster")
	flag.Usage = func() {
		info := "\t[*]App for searching and save info about movies from http://www.omdbapi.com/\n"
		info += "\t[*]Flag: -d (true|false) defaul false\n\tit's used for downloading movie poster\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		
	}
}

func useInterface() {
	flag.Parse()
	args = flag.Arg(0)
	om := poster.OmdAPI{}
	switch *d {
	case true:
		if err := om.Search(args, true); err != nil {
			fmt.Printf("Error  %s\n", err)
		}
	default:
		if err := om.Search(args, false); err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}

}

func main() {
	useInterface()

}
