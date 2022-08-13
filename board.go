package main

import "fmt"

type Tile string

const cellsPerSide = 3

// Enum of Tile values
const (
	empty  Tile = " "
	circle Tile = "O"
	cross  Tile = "X"
)

type Board struct {
	matrix [cellsPerSide][cellsPerSide]Tile
}

func (b *Board) create() {
	// Initialize playing board with empty tiles
	for i := 0; i < cellsPerSide; i++ {
		for j := 0; j < cellsPerSide; j++ {
			b.matrix[i][j] = empty
		}
	}
}

func (b *Board) print() {
	for i := 0; i < cellsPerSide; i++ {
		for j := 0; j < cellsPerSide; j++ {
			fmt.Printf("%s", b.matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func (b *Board) insert(position Coords, player Player) {
	b.matrix[position.row][position.col] = player.playerCharacter
}
