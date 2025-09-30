package main

import (
	"fmt"
	"log"
	"strings"
)

/*
 * CHESS KNOWLEDGE:
 * 1. Rank = row, File = col
 * 2. idk, will see
 *
 * TODO:
 * 1. Cell highlight based input in client
 * 2. Validation logic
 * 3. Advanced string parsing
 * 4. Movement logic
 */

const (
	WHITE = "\033[38;5;196m" // bright red
	BLACK = "\033[38;5;21m"  // bright blue
	RESET = "\033[0m"
	SIZE  = 9
)

type Move struct {
	From, To        Point
	Capture         bool
	Promote         bool
	PromotionChoice byte
	Piece           *Piece
	Castle          uint8
}

type Castle struct {
}

type Point struct{ R, C byte }

type Piece struct {
	Unit  byte
	Color bool
}
type Board struct {
	Grid     [][]*Piece
	Turn     bool
	Captured []byte
}

func (B *Board) Init() {
	B.Grid = make([][]*Piece, SIZE)
	initial := []string{
		"rnbqkbnr", // 8th, black
		"pppppppp",
		"........",
		"........",
		"........",
		"........",
		"PPPPPPPP",
		"RNBQKBNR", // 1st, white
	}
	for i := range len(B.Grid) {
		B.Grid[i] = make([]*Piece, SIZE)
		for j := range len(B.Grid[i]) {
			B.Grid[i][j] = &Piece{
				Unit:  initial[i][j],
				Color: B.Grid[i][j].Unit >= 65 && B.Grid[i][j].Unit <= 90,
			}
		}
	}
	B.Turn = true
	B.Captured = []byte{}
}

func (B *Board) IsValidPawnMove(m Move) bool {

}

func (B *Board) IsValidKnightMove(m Move) bool {

}

func (B *Board) IsValidRookMove(m Move) bool {

}

func (B *Board) IsValidMove(m Move) bool {
	piece := B.Grid[m.From.R][m.From.C]
	if piece == nil || piece.Color != B.Turn {
		return false
	}
	switch piece.Unit {

	}
	return false
}

func (B *Board) ParseMove(input string) {
	var move Move
	var ptr uint8 = 0

	switch strings.TrimSpace(input) {
	case "0-0":
		move.Castle = 1
		return
	case "0-0-0":
		move.Castle = 2
		return
	}

	if input[ptr] >= 97 && input[ptr] <= 104 {
		if input[ptr+1] == 'x' {
			move.Capture = true
			move.From.C = input[ptr]
			// find a way to get the row as well by searching the column for the pawn
			ptr += 2
		}

		if B.Turn {
			move.From = Point{}
			move.Piece.Unit = 'p'
		} else {
			move.Piece.Unit = 'P'
		}
	} else {
		move.Piece.Unit = input[ptr]
		ptr++
		if input[ptr] == 'x' {
			move.Capture = true
			ptr++
		}
	}
	// if move.SourceFile == 0 {
	// move.SourceFile = 'z'
	// }

	move.To.C = input[ptr]
	ptr++
	move.To.R = input[ptr]
	ptr++
	if int(ptr) < len(input) && input[ptr] == '=' {
		move.Promote = true
		ptr++
	}
	fmt.Printf("pos: %c%c | unit: %c | capture: %t | sourceFile: %c | promote: %t\n",
		move.To.C, move.To.R, move.Piece.Unit, move.Capture, move.From.C, move.Promote)

}

func (B *Board) Move(move string) {
	B.ParseMove(move)
	B.Turn = !B.Turn
}

func (B *Board) Print() {
	for _, r := range B.Grid {
		for _, el := range r {
			fmt.Printf("%s%c%s\t   ", WHITE, el, RESET)
		}
		fmt.Print("\n\n\n")
	}
}

func main() {
	var b Board
	b.Init()

	var move string
	for {
		_, err := fmt.Scanln(&move)
		if err != nil {
			log.Fatalln("error reading input:", err)
		}
		b.Move(move)
		b.Print()
	}
}
