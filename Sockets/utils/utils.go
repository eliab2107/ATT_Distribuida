package main

import (
    "fmt"
    "time"
)

const (
    width  = 4
	height = 4
)

var initialBoard = [][]bool	{
		{false, true, false, false},
		{true, false, true, false},
		{false, false, true, true },
		{false, false, false, false },
}

var initialNewBoard = [][]bool	{
    {false, false, false, false, false },
    {false, false, false, false, false },
    {false, false, false, true, false},
    {false, false, false, false, false},
    {false, false, false, false, false},
}
// Cria uma matriz bidimensional representando o tabuleiro
func makeBoard(width, height int) [][]bool {
    board := make([][]bool, height)
    for i := range board {
        board[i] = make([]bool, width)
    }
    return board
}

// Inicializa o tabuleiro com células vivas de forma aleatória
// Inicializa o tabuleiro com uma matriz fixa
func initializeBoard(board [][]bool, initialBoard [][]bool) {
    for i := range board {
        for j := range board[i] {
            board[i][j] = initialBoard[i][j]
        }
    }
}

// Conta o número de vizinhos vivos ao redor de uma célula
func countNeighbors(board [][]bool, x, y int) int {
    var count int
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if i == 0 && j == 0 {
                continue
            }
            nx, ny := x+i, y+j
            if nx >= 0 && nx < len(board[0]) && ny >= 0 && ny < len(board) && board[ny][nx] {
                count++
            }
        }
    }
    return count
}

// Atualiza o tabuleiro baseado nas regras do Jogo da Vida
func updateBoard(board [][]bool) [][]bool {
    newBoard := makeBoard(len(board[0]), len(board))
    for y := range board {
        for x := range board[y] {
            neighbors := countNeighbors(board, x, y)
            if board[y][x] {
                newBoard[y][x] = neighbors == 2 || neighbors == 3
            } else {
                newBoard[y][x] = neighbors == 3
            }
        }
    }
    return newBoard
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

func isEqual(board1, board2 [][]bool) bool {
    if len(board1) != len(board2) {
        return false
    }
    for i := range board1 {
        if len(board1[i]) != len(board2[i]) {
            return false
        }
        for j := range board1[i] {
            if board1[i][j] != board2[i][j] {
                return false
            }
        }
    }
    return true
}


func main() {
    board := makeBoard(width, height)
    initializeBoard(board, initialBoard)
    for i:=0; i<100; i++{
        printBoard(board)
        newBoard := updateBoard(board)
        if isEqual(board, newBoard) {
            fmt.Println("O jogo atingiu um estado estável.")
            break
        }
        board = newBoard
        time.Sleep(time.Second)
        fmt.Println()
    }
   
}