package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/template"
	"xcd/xcdstruct"
)

var fetch *bool

var temp = template.Must(template.New("XCDComic").Parse(`
	Month: {{.Month}} 
	NumID: {{.NumID}}  
	Year:  {{.Year}} 
	News: {{.News}}
	SafeTitle: {{.SafeTitle}} 
	Transcript: {{.Transcript}}  
	Image: {{.Image}} 
	Alt: {{.Alt}} 
`))

func init() {
	fetch = flag.Bool("fe", false, "fetch data on https://xkcd.com")
	flag.Usage = func() {
		info := "[*] Flags: -fetch  (true|false) defaul false\nit's used for downloading data"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func fetchingOnline() {
	fmt.Println("Fetching...")
	arrayData := xcdstruct.Comics{}
	for i := 1; i < 50; i++ {
		data, err := xcdstruct.SearchComic(i)
		if err != nil {
			fmt.Printf("%s fetching %d", err, i)
		}
		arrayData = append(arrayData, data)
	}
	if err := xcdstruct.Serializer(&arrayData); err != nil {
		fmt.Printf("Error saving data %s", err)
	}

}

func fetchingLocal(id int) {
	arrayData := xcdstruct.Comics{}
	if err := xcdstruct.Unserializer(&arrayData); err != nil {
		fmt.Printf("Error in unserialize data %s", err)
		return
	}

	for _, item := range arrayData {
		if item.NumID == id {
			temp.Execute(os.Stdout, item)
			return
		}
	}

}

func handlerAgs(arg string) int {
	id, err := strconv.Atoi(arg)
	if err != nil {
		panic("Argument is not a number")
	}
	return id
}

func useInterface() {
	flag.Parse()
	args := flag.Arg(0)
	switch *fetch {
	case true:
		fetchingOnline()
		fmt.Println("Successful download!")
	default:
		id := handlerAgs(args)
		fetchingLocal(id)

	}

}

func main() {
	useInterface()

}
