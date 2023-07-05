package main

import (
	"fmt"
)

func main() {
	countr := 0
	matrix := make([][]int, 10)
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			countr++
			matrix[x][y] = countr

		}
		fmt.Println(matrix[x])
	}

}
