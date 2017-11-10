package main

import (
"math/rand"
"time"
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

func set_mirror(board [][]int, id int) [][]int {
     level := id / 3
     col := id - 3*level
     //increment all vals in the row/col/diagonal where an
     //an entry was just played
     //need to not just increment but need to keep
     //track of how many in that line are filled
     for i := 0; i < 3; i++ {
     	 board[level][i] += 1
	 if i != level {
	    board[i][col] += 1
	    }
	 }
     if id % 2 == 0 {
     	//fill in diagonal
     	if level == col {
	   for i := 0; i < 3; i++ {
	       if i != level {
	       	  board[i][i] +=1
	       	  }
	       }
	       }
	if level != col {
	   for i:= 0; i < 3; i++ {
	    if i != level {
	       board[i][2 - i] += 1
	    }
	   }
	   }
	   }
     return board
}


func get_board_entry(board [][]int, id int) int {
     level := id / 3
     col := id - 3*level
     return board[level][col]
}

func ai_choose(board [][]int, mirror [][]int) int {
     //find lines with sum 2
     //find entries w/ board value == 0 (no one's played there yet)
     max_i := 0
     max_j := 0
     max_value := 0
     for i := range board {
     	 row := board[i]
     	 for j := range row {
	     if row[j] == 1 {
	     	continue
	     }
	     if  max_value < mirror[i][j] {
	     	 max_value = mirror[i][j]
		 max_i = i
		 max_j = j
		 }
	 }
     }
     return 3 * max_i + max_j
}

func ai_move(board [][]int, index int) [][]int{
     time.Sleep(time.Second * 1)
     fmt.Println("AI move")
     board = set_board(board, index, -1)
     return board
}

func main() {
     board := [][]int{{0,0,0}, {0,0,0}, {0,0,0}}
     mirror := [][]int{{0,0,0}, {0,0,0}, {0,0,0}}
     for {
     	 fmt.Print("Enter position (0-8): ")
     	 scanner := bufio.NewScanner(os.Stdin)
     	 for scanner.Scan() {
             lineStr := scanner.Text()
    	     num, _ := strconv.Atoi(lineStr)
	     mirror = set_mirror(mirror, num)
	     fmt.Println("Mirror")
	     print_board(mirror)
	     board = set_board(board, num, 1)
	     break
	 }
	 fmt.Println("Board")
	 print_board(board)
	 ai_num := ai_choose(board, mirror)
	 for get_board_entry(board, ai_num) != 0 {
	     ai_num = rand.Intn(8)
	 }
	 ai_move(board, ai_num)
	 print_board(board)
	 if check_board(board) == 1 {
	    break
	 }	     
    }	 
}