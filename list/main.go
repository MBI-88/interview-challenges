package main

import (
	"bufio"
	"fmt"
	"list/structs"
	"os"
	"strconv"
)

func userInterface() {
	var list *structs.List
	info := "\t\t**** Welcome to Linked list structure ****\n"
	info += "\t Options\n"
	info += "\t Add    (1)\n"
	info += "\t Print  (2)\n"
	info += "\t Delete (3)\n"
	info += "\t Find   (4)\n"
	info += "\t Size   (5)\n"
	info += "\t Clean  (6)\n"
	info += "\t Exit   (7)\n"

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
			fmt.Println("Append")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				list = structs.Append(list, data)
				goto read
			}
		case 2:
			fmt.Println("Print")
			structs.Print(list)
			fmt.Println()
			goto read
		case 3:
			fmt.Println("Delete")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				list = structs.Delete(list, data)
				goto read
			}
		case 4:
			fmt.Println("Find")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				fmt.Printf("Data to find %d data found %d\n", data, structs.Find(list, data))
				goto read
			}
		case 5:
			fmt.Println("Size")
			fmt.Printf("%d\n", structs.Len(list))
			goto read
		case 6:
			fmt.Println("Clean")
			list = structs.Clean(list)
			goto read
		case 7:
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
