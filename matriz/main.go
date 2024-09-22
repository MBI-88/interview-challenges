package main

import "fmt"

/*
	Problema llenar una matriz de forma espiral
*/

// Llena una matriz por medio de un array. Args: matriz,array,rowmax,colmax
// returna la matriz llena de forma espiral
func fillMatriz(matriz [3][4]int, array [12]int, rowmax, colmax int) [3][4]int {
	var (
		row int
		col int
	)
	step := 1
	leftstep := 0
	for p := 0; p < len(array); p++ {
		if step == 1 && matriz[row][col] == 0 { // left to right
			matriz[row][col] = array[p]
			col++
			if col == colmax-1 {
				step = 2
			}
		} else if step == 2 && matriz[row][col] == 0 { // top to buttom
			matriz[row][col] = array[p]
			row++
			if row == rowmax-1 {
				step = 3
			}
		} else if step == 3 && matriz[row][col] == 0 { // right to left
			matriz[row][col] = array[p]
			col--
			if col == leftstep {
				step = 4
			}

		} else if step == 4 && matriz[row][col] == 0 { // buttom to top
			matriz[row][col] = array[p]
			row--
			if row == rowmax-(rowmax-1) {
				step = 1
				colmax--
				rowmax--
				leftstep++
			}
			
			if array[p] == len(array) - 1 {
				col++
				row++
			}
			

		}

	}
	return matriz
}

func main() {
	var (
		matriz [3][4]int
	)
	array := [12]int{1, 2, 3, 4,5,6,7,8,9,10,11,12}
	matriz = fillMatriz(matriz, array, 3, 4)
	fmt.Println(matriz)

}
