package main

import (
	"flag"
	"fmt"
	"os"
	"unique/cmd"
)



var word *string

func init() {
	word = flag.String("word", "", "set word to test")
	flag.Usage = func() {
		info := fmt.Sprintln("[*] Unique character problem")
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func operation() {
	flag.Parse()
	fmt.Printf("%s is unique? %t\n", *word, cmd.UniqueCharacters(*word))

}


func main() {
	operation()

}