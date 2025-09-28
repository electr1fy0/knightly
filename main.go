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
	Pos        Point
	Capture    bool
	Promote    bool
	Unit       byte
	SourceFile byte
	Castle     uint8
}

type Castle struct {
}

type Point struct{ R, C byte }

type Board struct {
	Grid     [][]byte
	Turn     bool
	Captured []byte
}

func (B *Board) Init() {
	B.Grid = make([][]byte, SIZE)
	initial := "8rnbqkbnr7pppppppp6........5........4........3........2PPPPPPPP1RNBQKBNR abcdefgh"
	for i := range len(B.Grid) {
		B.Grid[i] = make([]byte, SIZE)
		for j := range len(B.Grid[i]) {
			B.Grid[i][j] = initial[i*SIZE+j]
		}
	}
}

func (B *Board) ParseMove(input string) {
	var move Move
	var ptr uint8 = 0

	switch strings.TrimSpace(input) {
	case "0-0":
		move.Castle = 1
	case "0-0-0":
		move.Castle = 2
	}

	if input[ptr] >= 97 && input[ptr] <= 104 {
		if input[ptr+1] == 'x' {
			move.Capture = true
			move.SourceFile = input[ptr]
			ptr += 2
		}
		if B.Turn {
			move.Unit = 'p'
		} else {
			move.Unit = 'P'
		}
	} else {
		move.Unit = input[ptr]
		ptr++
		if input[ptr] == 'x' {
			move.Capture = true
			ptr++
		}
	}
	if move.SourceFile == 0 {
		move.SourceFile = 'z'
	}

	move.Pos.C = input[ptr]
	ptr++
	move.Pos.R = input[ptr]
	ptr++
	if int(ptr) < len(input) && input[ptr] == '=' {
		move.Promote = true
		ptr++
	}
	fmt.Printf("pos: %c%c | unit: %c | capture: %t | sourceFile: %c | promote: %t\n",
		move.Pos.C, move.Pos.R, move.Unit, move.Capture, move.SourceFile, move.Promote)

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
