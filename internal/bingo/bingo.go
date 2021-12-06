package bingo

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
	return Game{}
}

func (g *Game) Play() Board {
	return Board{}
}
