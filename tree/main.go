package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tree/structs"
)

func userInterface() {
	var tree *structs.Tree 
	info := "\t\t**** Welcome to Tree structure ****\n"
	info += "\t Options\n"
	info += "\t Add (1)\n"
	info += "\t Print (pre order) (2)\n"
	info += "\t Print (pos order) (3)\n"
	info += "\t Print (in order)  (4)\n"
	info += "\t Delete            (5)\n"
	info += "\t Find              (6)\n"
	info += "\t Total Leave       (7)\n"
	info += "\t Check tree        (8)\n"
	info += "\t Show tree         (9)\n"
	info += "\t Depht tree        (10)\n"
	info += "\t Exit              (11)\n"

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
			fmt.Println("Add number")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				tree = structs.Append(tree, data)
				goto read
			}

		case 2:
			fmt.Println("Print tree (pre order)")
			structs.PreOrder(tree)
			fmt.Println()
			goto read
		case 3:
			fmt.Println("Print tree (pos order)")
			structs.PosOrder(tree)
			fmt.Println()
			goto read
		case 4:
			fmt.Println("Print tree (in order)")
			structs.InOrder(tree)
			fmt.Println()
			goto read
		case 5:
			fmt.Println("Delete")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				structs.DelNode(tree, data)
				fmt.Println()
				goto read
			}
		case 6:
			fmt.Println("Find")
			b := bufio.NewScanner(os.Stdin)
			for b.Scan() {
				data, err := strconv.Atoi(b.Text())
				if err != nil {
					fmt.Println("Data must be a number")
					goto read
				}
				fmt.Printf("Data to search %d data found %d\n", data, structs.FindNode(tree, data))
				goto read
			}
		case 7:
			fmt.Println("Total Leave")
			fmt.Printf("Total leave %d\n", structs.Leave(tree))
			goto read 
		case 8:
			fmt.Println("Check tree")
			fmt.Printf("Tree is empty %v\n", structs.IsEmpty(tree))
			goto read
		case 9:
			fmt.Println("Show tree")
			structs.ShowTree(tree,0)
			goto read
		case 10:
			fmt.Println("Depht tree")
			fmt.Printf("Depht %d\n",structs.Depth(tree))
			goto read
		case 11:
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
