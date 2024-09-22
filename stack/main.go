package main

import (
	"bufio"
	"fmt"
	"os"
	"stack/structs"
	"strconv"
)

func userInterface() {
	var stack *structs.Stack
	info := "\t\t**** Welcome to Stack structure ****\n"
	info += "\t Options\n"
	info += "\t Add   (1)\n"
	info += "\t Print (2)\n"
	info += "\t Pop   (3)\n"
	info += "\t Find  (4)\n"
	info += "\t Size  (5)\n"
	info += "\t Exit  (6)\n"

	buffer := bufio.NewScanner(os.Stdin)
read:
	fmt.Println(info)
	for buffer.Scan() {
		op, err := strconv.Atoi(buffer.Text())
		if err != nil {
			fmt.Printf("Error %s", err)
			goto read
		}
		switch op {
		case 1:
			fmt.Println("Push")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				stack = structs.Push(stack, data)
				goto read
			}
		case 2:
			fmt.Println("Print")
			structs.Print(stack)
			fmt.Println()
			goto read
		case 3:
			fmt.Println("Pop")
			stack = structs.Pop(stack)
			fmt.Println()
			goto read
		case 4:
			fmt.Println("Find")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				fmt.Printf("Data to find %d data found %d\n", data, structs.Find(stack, data))
				goto read
			}
		case 5:
			fmt.Println("Size")
			fmt.Printf("%d\n", structs.Len(stack))
			goto read
		case 6:
			break
		default:
			goto read

		}
		fmt.Println("Exit...")
		os.Exit(0)
	}
}

func main() {
	userInterface()
}
