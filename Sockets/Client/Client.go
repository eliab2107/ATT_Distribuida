package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

const (
    width  = 4
	height = 4
)

var initialBoard = [][]bool	{
		{false, true, false, false, false },
		{true, false, true, false, false },
		{false, false, true, true, false},
		{false, false, false, false, false},
		{false, false, true, false, true},
}

// Cria uma matriz bidimensional representando o tabuleiro
func makeBoard(width, height int) [][]bool {
    board := make([][]bool, height)
    for i := range board {
        board[i] = make([]bool, width)
    }
    return board
}

// Inicializa o tabuleiro com uma matriz fixa
func initializeBoard(board [][]bool, initialBoard [][]bool) {
    for i := range board {
        for j := range board[i] {
            board[i][j] = initialBoard[i][j]
        }
    }
}





// Imprime o tabuleiro no console
func printBoard(board [][]bool) {
    for _, row := range board {
        for _, cell := range row {
            if cell {
                fmt.Print("O ")
            } else {
                fmt.Print(". ")
            }
        }
        fmt.Println()
    }
}

func main() {
	// Definindo as variaveis
	board := makeBoard(width, height)
    initializeBoard(board, initialBoard)
	response := make([]byte, 4096)
	var newMatriz [][]bool
	conn, err := net.Dial("tcp", "127.0.0.1:1313")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Serializa a matriz para JSON
	data, err := json.Marshal(board)
	if err != nil {
		log.Fatal(err)
	}

	// Envia a matriz serializada pela conexão
	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	for {
		n, err := conn.Read(response)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(response[:n], &newMatriz)
		printBoard(newMatriz)
		fmt.Println()
		if err != nil {
			fmt.Println("Estado final alcançado: ")
			break
		}
		
	}
	
}
