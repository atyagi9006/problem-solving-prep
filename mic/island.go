package main

import "fmt"

func main() {
	il := &island{
		Grid: [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0}},
	}

	fmt.Println("Parameter of island", il.parameter())

}

type island struct {
	Grid [][]int
}

func (il *island) parameter() int {
	land := 0
	edge := 0
	for i := 0; i < len(il.Grid); i++ {
		for j := 0; j < len(il.Grid[i]); j++ {
			if il.Grid[i][j] == 1 {
				land++
				if i > 0 && il.Grid[i-1][j] == 1 {
					edge++
				}
				if j > 0 && il.Grid[i][j-1] == 1 {
					edge++
				}
			}
		}
	}
	res := land*4 - edge*2
	return res
}
