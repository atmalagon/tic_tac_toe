package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func boardColumn(board [][]int, columnIndex int) []int {
	column := make([]int, 3)
	for i, row := range board {
		column[i] = row[columnIndex]
	}
	return column
}

func calculate_sum(line []int) (sum int) {
	sum = 0
	for _, h := range line {
		sum += h
	}
	return
}

func get_score(board [][]int) int {
	//check rows and columns
	for i := 0; i < 3; i++ {
		col := boardColumn(board, i)
		col_sum := calculate_sum(col)
		row_sum := calculate_sum(board[i])
		if Abs(col_sum) == 3 {
			return col_sum
		}
		if Abs(row_sum) == 3 {
			return row_sum
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

func game_state(board [][]int, id int, val int) [][]int {
	level := id / 3
	col := id - 3*level
	duplicate := make([][]int, len(board))
	for i := range board {
		duplicate[i] = make([]int, len(board[i]))
		copy(duplicate[i], board[i])
	}
	duplicate[level][col] = val
	return duplicate
}

func get_board_entry(board [][]int, id int) int {
	level := id / 3
	col := id - 3*level
	return board[level][col]
}

func possible_moves(board [][]int) []int {
	//return list of open positions
	my_list := make([]int, 0)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				my_list = append(my_list, 3*i+j)
			}
		}
	}
	return my_list
}

func minimax(board [][]int, player int) (move, score int) {
	// if game is over, assign win points
	score = get_score(board)
	if Abs(score) == 3 {
		return -1, score
	}
	//otherwise get list of new game states
	open_positions := possible_moves(board)
	if len(open_positions) == 0 {
		return -1, -1
	}
	scores := make([]int, len(open_positions))
	for i := 0; i < len(open_positions); i++ {
		state := game_state(board, open_positions[i], player)
		player2 := -1 * player
		_, state_score := minimax(state, player2)
		scores[i] = state_score
	}
	//if player=1, get move that maximizes score, otherwise find max
	if player == 1 {
		best_move := open_positions[0]
		max_score := scores[0]
		for i := 1; i < len(scores); i++ {
			if scores[i] > max_score {
				max_score = scores[i]
				best_move = best_move
			}
		}
		return best_move, max_score
	} else if player == -1 {
		best_move := open_positions[0]
		min_score := scores[0]
		for i := 1; i < len(scores); i++ {
			if scores[i] < min_score {
				min_score = scores[i]
				best_move = best_move
			}
		}
		return best_move, min_score
	}
	return
}

func ai_move(board [][]int, index int) [][]int {
	//time.Sleep(time.Second * 1)
	fmt.Println("AI move")
	board = game_state(board, index, -1)
	return board
}

func main() {
	board := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for {
		fmt.Print("Enter position (0-8): ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lineStr := scanner.Text()
			num, _ := strconv.Atoi(lineStr)
			board = game_state(board, num, 1)
			break
		}
		fmt.Println("Board")
		print_board(board)
		ai_num, score := minimax(board, -1)
		if ai_num == -1 && score != -1 {
			fmt.Println("Tic Tac Toe")
			break
		} else if ai_num == -1 && score == -1 {
			fmt.Println("Stalemate")
			break
		}
		board = ai_move(board, ai_num)
		print_board(board)
	}
}
