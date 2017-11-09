package main

import (
"math/rand"
"bufio"
"os"
"fmt"
"strconv"
)

func boardColumn(board [][]int, columnIndex int) []int {
    column := make([]int, 3)
    for i, row := range board {
        column[i] = row[columnIndex]
    }
    return column
}

func calculate_sum(line []int) int {
     sum := 0
     for _, h := range line {
     	 sum += h
     }
     if sum == 3 {
     	fmt.Println("Tic Tac Toe!")
	return 1
     }
     return 0
}

func check_board(board [][]int) int {
     //check rows and columns
     for i := 0; i < 3; i++ {
     	 col := boardColumn(board, i)
	 if (calculate_sum(col) == 1) {
	    return 1
	 }
	 if (calculate_sum(board[i]) == 1) {
     	    return 1
     	 }
     }
     //check diagonals
     left := board[0][0] + board[1][1] + board[2][2]
     right := board[0][2] + board[1][1] + board[2][0]
     if (left == 3) {
     	fmt.Println("Tic Tac Toe!")
     	return 1
     }
     if  (right == 3) {
     	fmt.Println("Tic Tac Toe!")
     	return 1
     }
     return 0
}

func print_board(board [][]int) {
     fmt.Println(board[0])
     fmt.Println(board[1])
     fmt.Println(board[2])

}

func set_board(board [][]int, id int, val int) [][]int {
     level := id / 3
     col := id - 3*level
     board[level][col] = val
     return board
}

func ai_move(board [][]int, index int) [][]int{
     board = set_board(board, index, 2)
     return board
}

func main() {
     board := [][]int{{0,0,0}, {0,0,0}, {0,0,0}}
     for {
     	 fmt.Print("Enter position (0-8): ")
     	 scanner := bufio.NewScanner(os.Stdin)
     	 for scanner.Scan() {
             lineStr := scanner.Text()
    	     num, _ := strconv.Atoi(lineStr)
	     board = set_board(board, num, 1)
	     break
	 }
	 print_board(board)
	 fmt.Println("AI move")
	 ai_move(board, rand.Intn(8))
	 if check_board(board) == 1 {
	    break
	 }	     
    }	 
}