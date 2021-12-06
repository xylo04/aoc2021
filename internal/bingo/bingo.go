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
			board = Board{NumberGrid: make([][]uint, 5)}
		}
		if y == 0 {
			board.NumberGrid[x] = make([]uint, 5)
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
	return Board{}
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
