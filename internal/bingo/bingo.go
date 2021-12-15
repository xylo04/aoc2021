package bingo

import (
	"strconv"
	"strings"
	"text/scanner"
)

type Board struct {
	NumberGrid  [][]uint
	Marked      [][]bool
	WinsInRound uint
	Score       uint
}

type Game struct {
	Balls  []uint
	Boards []Board
}

func ParseGame(in string) Game {
	game := Game{}
	var s scanner.Scanner
	var tok rune
	s.Init(strings.NewReader(in))
	s.Filename = "game input"
	for tok = s.Scan(); tok != scanner.EOF && s.Line < 3; tok = s.Scan() {
		if s.TokenText() != "," {
			b, _ := strconv.ParseUint(s.TokenText(), 10, 32)
			game.Balls = append(game.Balls, uint(b))
		}
	}
	var x, y = uint(0), uint(0)
	var board Board
	for ; tok != scanner.EOF; tok = s.Scan() {
		if x == 0 && y == 0 {
			board = Board{NumberGrid: make([][]uint, 5), Marked: make([][]bool, 5)}
		}
		if y == 0 {
			board.NumberGrid[x] = make([]uint, 5)
			board.Marked[x] = make([]bool, 5)
		}
		n, _ := strconv.ParseUint(s.TokenText(), 10, 32)
		board.NumberGrid[y][x] = uint(n)
		x = (x + 1) % 5
		if x == 0 {
			y = (y + 1) % 5
		}
		if x == 0 && y == 0 {
			game.Boards = append(game.Boards, board)
		}
	}
	return game
}

func (g *Game) Play() Board {
	for i, ball := range g.Balls {
		for _, board := range g.Boards {
			win := board.Turn(ball)
			if win {
				board.WinsInRound = uint(i + 1)
				board.TallyScore(ball)
				return board
			}
		}
	}
	return Board{}
}

func (b *Board) Turn(ball uint) bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.NumberGrid[y][x] == ball {
				b.Marked[y][x] = true
				return b.CheckWin()
			}
		}
	}
	return false
}

func (b *Board) CheckWin() bool {
	// check rows
	for y := 0; y < 5; y++ {
		if b.Marked[y][0] && b.Marked[y][1] && b.Marked[y][2] && b.Marked[y][3] && b.Marked[y][4] {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		if b.Marked[0][x] && b.Marked[1][x] && b.Marked[2][x] && b.Marked[3][x] && b.Marked[4][x] {
			return true
		}
	}
	return false
}

func (b *Board) String() string {
	var buf = ""
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			buf += " " + strconv.Itoa(int(b.NumberGrid[y][x]))
		}
		buf += "\n"
	}
	return buf
}

func (b *Board) TallyScore(ball uint) {
	var sum uint
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.Marked[y][x] {
				sum += b.NumberGrid[y][x]
			}
		}
	}
	b.Score = sum * ball
}
