package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"time"
)

const (
	width  = 5
	height = 5
)

var initialBoard = [][]bool{
	{false, true, false, false, false},
	{true, false, true, false, false},
	{false, false, true, true, false},
	{false, false, false, false, false},
	{false, false, true, false, true},
}

func calculateMedia(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, duration := range durations {
		total += duration
	}
	average := total / time.Duration(len(durations))
	return average
}

func calculateStandardDeviation(durations []time.Duration, mean time.Duration) float64 {
	var variance float64
	meanMs := float64(mean.Nanoseconds()) / 1e6 // Converter a média para milissegundos
	for _, duration := range durations {
		durationMs := float64(duration.Nanoseconds()) / 1e6 // Converter duração para milissegundos
		diff := durationMs - meanMs
		variance += diff * diff
	}
	variance /= float64(len(durations))
	return math.Sqrt(variance)
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
	// Definindo as variáveis
	board := makeBoard(width, height)
	initializeBoard(board, initialBoard)
	response := make([]byte, 1024)
	var newMatriz [][]bool
	var mediaList []time.Duration
	var repeticoes = 100
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
	for i := 0; i < repeticoes; i++ {
		timeatual := time.Now()
		_, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		for {
			n, err := conn.Read(response)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(response[:n], &newMatriz)
			if err != nil {
				break
			}
		}
		mediaList = append(mediaList, time.Since(timeatual))
	}
	tempoMedio := calculateMedia(mediaList)
	desvioPadrao := calculateStandardDeviation(mediaList, tempoMedio)

	// Converte tempo médio e desvio padrão para milissegundos
	tempoMedioMs := float64(tempoMedio.Nanoseconds()) / 1e6 // Convertendo de nanosegundos para milissegundos
	desvioPadraoMs := desvioPadrao // Desvio padrão já está em milissegundos

	fmt.Printf("Repetições: %d, Tempo total: %v, Tempo médio: %.2f ms, Desvio padrão: %.2f ms\n", repeticoes, time.Since(timestarted), tempoMedioMs, desvioPadraoMs)
}
