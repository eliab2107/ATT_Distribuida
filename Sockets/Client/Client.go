package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const (
    width  = 3
	height = 3
)
var initialBoard = [][]bool	{
    {false, true, false},
    {true, false, false,},
    {false, false, false },
}

func calculateMedia(durations []time.Duration) time.Duration {
    var total time.Duration
    for _, duration := range durations {
        total += duration
    }
    average := total / time.Duration(len(durations))
    return average
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
	response := make([]byte, 96)
	
	var mediaList []time.Duration
	var repeticoes = 5000
	cont := 0
	conn, err := net.Dial("tcp", "127.0.0.1:1313")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Serializa a matriz para JSON
	data, err := json.Marshal(board)
	if err != nil {
		log.Fatal(err)
	}
	timestarted := time.Now()
	// Envia a matriz serializada pela conexão
	for i:=0;i<repeticoes;i++{
		timeatual := time.Now()
		_, _ = conn.Write(data)

		for j:=0;j<=1;j++{

			_, _ = conn.Read(response)
				
		}
		cont++
		mediaList = append(mediaList, time.Now().Sub(timeatual))
	}
	tempoMedio := calculateMedia(mediaList)
	fmt.Println("Repetições: ", repeticoes, "Tempo total: ", time.Now().Sub(timestarted), "Tempo médio: ", tempoMedio)
}
