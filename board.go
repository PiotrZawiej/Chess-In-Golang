package main

import "fmt"

type boardStructure struct {
	board [8][8]string
}

func (b *boardStructure) drawBoard() [8][8]string {

	for raw := range b.board {
		for field := range b.board[raw] {
			b.board[raw][field] = "[]"
			}
		}
	return b.board
}


func (b *boardStructure) dispalyBoard() {
	for raws := range b.drawBoard() {
		if raws == 0{
			fmt.Println(" A B C D E F G H")

		}

		for fields := range(b.board[raws]){
			fmt.Print(b.drawBoard()[raws][fields])
		}
		fmt.Println(raws + 1)
	}
}
