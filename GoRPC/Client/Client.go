package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

const (
    width  = 5
	height = 5
)
type GameOfLife struct{
	Board [][]bool
}
var initialBoard = [][]bool	{
	{false, true, false, false, false },
	{true, false, true, false, false },
	{false, false, true, true, false},
	{false, false, false, false, false},
	{false, false, true, false, true},
}



func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1313")
	if err != nil{
		log.Fatal("Erro ao conectar ao servidor RPC:", err)
	}

	defer client.Close()

	game := GameOfLife{
		Board: initialBoard,
	}

	contagem := time.Now()
	

	for i := 0; i < 1; i++{
		err = client.Call("GameOfLife.InitializeBoard", struct{}{}, &game.Board)
		if err != nil {
			log.Fatal("Erro ao inicializar o tabuleiro: ", err)
		}
	
		for {
			var reply string
			err = client.Call("GameOfLife.UpdateBoard", game.Board, &reply)
			if err != nil{
				log.Fatal("Erro ao atualizar a matriz: ", err)
			}
	
			if reply == "O jogo atingiu um estado estável."{
				
				break
			}
	
			err = json.Unmarshal([] byte(reply), &game.Board)
			printBoard(game.Board)
			if err != nil{
				log.Fatal("Erro ao decodificar o tabuleiro: ", err)
			}

			
		}

	}

	fmt.Println(time.Now().Sub(contagem))
	


}

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
	
