package main

import (
	"reflect"
	"testing"
)

var baseBoard = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

func Test_game_state(t *testing.T) {

	type args struct {
		board [][]int
		id    int
		val   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "place_1",
			args: args{
				board: baseBoard,
				id:    0,
				val:   1,
			},
			want: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "place -1",
			args: args{
				board: baseBoard,
				id:    3,
				val:   -1,
			},
			want: [][]int{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game_state(tt.args.board, tt.args.id, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("game_state() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimax(t *testing.T) {
	type args struct {
		board  [][]int
		player int
	}
	tests := []struct {
		name      string
		args      args
		wantMove  int
		wantScore int
	}{
		{
			name: "obvious row win",
			args: args{
				board:  [][]int{{-1, -1, 0}, {1, 1, 0}, {1, 0, 0}},
				player: -1,
			},
			wantMove:  2,
			wantScore: -3,
		},
		{
			name: "obvious diag win",
			args: args{
				board:  [][]int{{-1, 1, 0}, {1, -1, 0}, {1, 0, 0}},
				player: -1,
			},
			wantMove:  8,
			wantScore: -3,
		},
		{
			name: "initial move",
			args: args{
				board:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
				player: -1,
			},
			wantMove:  0,
			wantScore: -1,
		},
		{
			name: "initial move",
			args: args{
				board:  [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
				player: -1,
			},
			wantMove:  4,
			wantScore: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMove, gotScore := minimax(tt.args.board, tt.args.player)
			if gotMove != tt.wantMove {
				t.Errorf("minimax() gotMove = %v, want %v", gotMove, tt.wantMove)
			}
			if gotScore != tt.wantScore {
				t.Errorf("minimax() gotScore = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
