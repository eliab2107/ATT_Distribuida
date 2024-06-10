package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	//"time"
	"sync"
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

var initialNewBoard = [][]bool	{
    {false, false, false,},
    {false, false, false,},
    {false, false, false,},   
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

func conecta(server net.Conn, board [][]bool, wg *sync.WaitGroup){
	newBoard := makeBoard(width, height)
    
	for {
        newBoard = updateBoard(board)
		
		if isEqual(board, newBoard) {
			_, _ = server.Write([]byte("O jogo atingiu um estado estável."))
			break
		}
		
		data, err := json.Marshal(newBoard)
		if err != nil {
			log.Println("Erro ao enviar dados:", err)
			return
		}		
		  _, _ = server.Write(data)  
		board = newBoard 
        time.Sleep(500 * time.Microsecond)
	} 
    wg.Done()
}

func main(){
	r, _ :=net.ResolveTCPAddr("tcp","127.0.0.1:1313" )
	server, _ := net.ListenTCP("tcp", r)
	fmt.Println("Listen on IP and port: 127.0.0.1:1313")
    var wg sync.WaitGroup
	data := make([]byte, 96) // Tamanho do buffer de recebimento
	//Aceitando conexão
	conn, _ := server.Accept()
    fmt.Println("Conexao iniciada")
    for i:= 0; i<5000; i++{
		//Recebendo a a matriz padrão de entrada
		n, _ := conn.Read(data)
		var board [][]bool
		//Desserializando ela
		_ = json.Unmarshal(data[:n], &board)
		//Chamando a execução do jogo
        wg.Add(1)
		go conecta(conn, board, &wg)
	}
}