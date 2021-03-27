package main

import (
	"fmt"
)

func main() {
	fmt.Println("spirak Matrix: ", genMatrix(3))

}
func genMatrix(n int) [][]int {
	if n < 0 {
		return nil
	}
	eor := n
	eoc := n
	row := 0    //start of the row
	column := 0 //start of col
	counter := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for row < eor && column < eoc {
		//fill first row
		for i := column; i <eoc ; i++ {
			res[row][i] = counter
			counter++
		}
		fmt.Println("row filled: ",res)
		row++
		//fill the last column
		for i := row; i < eor; i++ {
			res[i][eoc-1] = counter
			counter++
		}
		fmt.Println("column filled:",res)
		eoc--

		//fill the lastrow from the remating rows
		if row < eor {
			for i := eoc - 1; i >= column; i-- {
				res[eor-1][i] = counter
				counter++
			}
			eor--
		}
		fmt.Println("other row",res)

		//fill the first column from the remaining columns
		if column < eoc {
			for i := eor - 1; i >= row; i-- {
				res[i][column] = counter
				counter++
			}
			column++
		}
		fmt.Println("other column :",res)
	}
	return res
}

func newGen(r int) [][]int {

	// Initialize value to be filled in matrix
	val := 1
	a := make([][]int, r)
	for i := 0; i < r; i++ {
		a[i] = make([]int, r)
	}
	/*  k - starting row index
	    m - ending row index
	    l - starting column index
	    n - ending column index */
	k := 0
	l := 0
	m := r
	n := r
	for k < m && l < n {
		/* Print the first row from the remaining
		   rows */
		for i := l; i < n; i++ {
			a[k][i] = val
			val++
		}

		k++

		/* Print the last column from the remaining
		   columns */
		for i := k; i < m; i++ {
			a[i][n-1] = val
			val++

		}
		n--

		/* Print the last row from the remaining
		   rows */
		if k < m {
			for i := n - 1; i >= l; i-- {
				a[m-1][i] = val
				val++
			}
			m--
		}

		/* Print the first column from the remaining
		   columns */
		if l < n {
			for i := m - 1; i >= k; i-- {
				a[i][l] = val
				val++
			}
			l++
		}
	}

	return a
}
