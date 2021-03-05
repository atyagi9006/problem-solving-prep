package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int{
		{0, 0, 1},
		{1, 0, 0},
		{1, 1, 0},
	}
	findWinner(a)

	a=[][]int{
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 0},
	}
	findWinner(a)
	
	a=[][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 0, 1},
	}
	findWinner(a)
	
	a=[][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 1, 1},
	}
	findWinner(a)
}

func findWinner(game [][]int) {
	
	rows := len(game)
	cols := len(game[0])

	if rows != cols {
		fmt.Println("invalid input")
		return
	}

	result := checkRows(game, rows)

	if result == 0 {
		result = checkColumns(game, rows)
	}

	if result == 0 {
		result = checkDiagonal(game, rows)
	}

	if result == 0 {
		result = checkOtherDiagonal(game, rows)
	}

	printResult(result)
}

func checkRows(game [][]int, limit int) int {
	//check rows
	for row := 0; row < limit; row++ {
		count := 0
		for col := 0; col < limit; col++ {
			if game[row][col] == 0 {
				count += -1
			} else {
				count++
			}
		}
		if count == limit || count == -limit {
			return count / int(math.Abs(float64(count)))
		}
	}
	return 0
}

func checkColumns(game [][]int, limit int) int {
	//check columns
	for col := 0; col < limit; col++ {
		count := 0
		for row := 0; row < limit; row++ {
			if game[row][col] == 0 {
				count += -1
			} else {
				count++
			}
		}
		if count == limit || count == -limit {
			return count / int(math.Abs(float64(count)))
		}
	}
	return 0
}

func checkDiagonal(game [][]int, limit int) int {
	//check diagonal start from game[0][0]
	count := 0
	for i := 0; i < limit; i++ {
		if game[i][i] == 0 {
			count += -1
		} else {
			count++
		}
	}
	if count == limit || count == -limit {
		return count / int(math.Abs(float64(count)))
	}
	return 0
}

func checkOtherDiagonal(game [][]int, limit int) int {
	//check other diagonal
	count2 := 0
	for i := 0; i < limit; i++ {
		if game[i][2-i] == 0 {
			count2 += -1
		} else {
			count2++
		}
	}
	if count2 == limit || count2 == -limit {
		return count2 / int(math.Abs(float64(count2)))
	}

	return 0
}

func printResult(result int) {
	switch result {
	case -1:
		fmt.Println("O wins")
		break
	case 1:
		fmt.Println("x wins")
		break
	default:
		fmt.Println("NO wins")
		break

	}
}
