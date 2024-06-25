package main

import (
	"context"
	"log"
	"time"
 	pb "path/to/your/proto/package" // Substitua pelo caminho correto para o pacote gerado
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGameOfLifeClient(conn)

	board := makeInitialBoard()
	req := &pb.BoardRequest{Rows: makeProtoFromBoard(board)}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.UpdateBoard(ctx, req)
	if err != nil {
		log.Fatalf("could not update board: %v", err)
	}
	newBoard := makeBoardFromProto(res)

	// Imprima ou utilize o novo tabuleiro (newBoard) conforme necessário
	log.Printf("Updated Board: %v", newBoard)
}

func makeInitialBoard() [][]bool {
	return [][]bool{
		{false, true, false, false, false },
		{true, false, true, false, false },
		{false, false, true, true, false},
		{false, false, false, false, false},
		{false, false, true, false, true},
	}
}

func makeProtoFromBoard(board [][]bool) []*pb.Board {
	height := len(board)
	protoBoard := make([]*pb.Board, height)
	for i, row := range board {
		protoRow := &pb.Board{Row: make([]*pb.Cell, len(row))}
		for j, cell := range row {
			protoRow.Row[j] = &pb.Cell{Alive: cell}
		}
		protoBoard[i] = protoRow
	}
	return protoBoard
}

func makeBoardFromProto(res *pb.BoardResponse) [][]bool {
	height := len(res.Rows)
	width := len(res.Rows[0].Row)
	board := make([][]bool, height)
	for i, row := range res.Rows {
		board[i] = make([]bool, width)
		for j, cell := range row.Row {
			board[i][j] = cell.Alive
		}
	}
	return board
}
