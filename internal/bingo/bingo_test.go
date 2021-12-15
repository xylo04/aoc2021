package bingo

import (
	"reflect"
	"testing"
)

func TestParseGame(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "website",
			args: args{in: `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`},
			want: Game{
				Balls: []uint{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				Boards: []Board{
					{
						NumberGrid: [][]uint{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
					{
						NumberGrid: [][]uint{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
					{
						NumberGrid: [][]uint{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseGame(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_Play(t *testing.T) {
	type fields struct {
		Balls  []uint
		Boards []Board
	}
	tests := []struct {
		name   string
		fields fields
		want   Board
	}{
		{
			name: "website",
			fields: fields{
				Balls: []uint{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				Boards: []Board{
					{
						NumberGrid: [][]uint{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
					{
						NumberGrid: [][]uint{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
					{
						NumberGrid: [][]uint{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
						Marked: [][]bool{
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
							{false, false, false, false, false},
						},
					},
				},
			},
			want: Board{
				NumberGrid: [][]uint{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Marked: [][]bool{
					{true, true, true, true, true},
					{false, false, false, true, false},
					{false, false, true, false, false},
					{false, true, false, false, true},
					{true, true, false, false, true},
				},
				WinsInRound: 12,
				Score:       4512,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Balls:  tt.fields.Balls,
				Boards: tt.fields.Boards,
			}
			if got := g.Play(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Play() = %v, want %v", got, tt.want)
			}
		})
	}
}
