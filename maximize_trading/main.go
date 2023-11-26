package main

import (
	"fmt"
	"maximize/cmd"
	"os"
)



func operation() {
	input := os.Args[1:]
	arrayInt, err := cmd.StringToInt(input)
	if err != nil {
		panic(err)
	}
	result := cmd.MaxProfitDays(arrayInt)
	fmt.Printf("To buy %d to sell %d\n",result[0],result[1])
	
}


func main() {
	operation()
}