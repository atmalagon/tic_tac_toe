package main

import (
"math/rand"
"math/Abs"
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
}

func score(board [][]int) int {
     //check rows and columns
     for i := 0; i < 3; i++ {
     	 col := boardColumn(board, i)
	 col_sum := calculate_sum(col)
	 row_sum := calculate_sum(board[i])
	 if Abs(col_sum) == 3 {
	    return col_sum
	 }
	 if Abs(board_sum) == 3 {
     	    return board_sum
     	 }
     }
     //check diagonals
     left := board[0][0] + board[1][1] + board[2][2]
     right := board[0][2] + board[1][1] + board[2][0]
     if Abs(left) == 3 {
     	return left
     }
     if Abs(right) == 3 {
     	return right
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
 
func get_board_entry(board [][]int, id int) int {
     level := id / 3
     col := id - 3*level
     return board[level][col]
}

func possible_moves(board) int {
     //return list of open positions
}

func minimax(board [][]int) int {
     // if game is over, return score
     score := score(board)
     if Abs(score) == 3 {
     	return score(board)  
     }
     //otherwise get list of new game states
     open_positions := possible_moves(board)
     scores = make([]int, 0)    
     for i := 0; i<len(open_positions); i++ {
     	 state := copy(board)
     	 state = set_board_entry(board, i)
     	 state_score = minimax(state)
	 scores = scores.append(scores, state_score)

     }


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
     moves := make([]int, 0)
     for {
     	 fmt.Print("Enter position (0-8): ")
     	 scanner := bufio.NewScanner(os.Stdin)
     	 for scanner.Scan() {
             lineStr := scanner.Text()
    	     num, _ := strconv.Atoi(lineStr)
	     moves = append(open_positions, num)
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