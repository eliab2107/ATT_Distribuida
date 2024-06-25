package main

import (
	"context"
	"log"
	"net"

	pb "path/to/your/proto/package" // Substitua pelo caminho correto para o pacote gerado

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Servidor implementa gameoflife.GameOfLifeServer
type server struct {
	pb.UnimplementedGameOfLifeServer
}

func (s *server) UpdateBoard(ctx context.Context, req *pb.BoardRequest) (*pb.BoardResponse, error) {
	board := makeBoardFromProto(req)
	newBoard := updateBoard(board)
	return &pb.BoardResponse{Rows: makeProtoFromBoard(newBoard)}, nil
}

func makeBoardFromProto(req *pb.BoardRequest) [][]bool {
	height := len(req.Rows)
	width := len(req.Rows[0].Row)
	board := make([][]bool, height)
	for i, row := range req.Rows {
		board[i] = make([]bool, width)
		for j, cell := range row.Row {
			board[i][j] = cell.Alive
		}
	}
	return board
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

func updateBoard(board [][]bool) [][]bool {
	// Implemente a lógica de atualização do Jogo da Vida aqui
	return board
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameOfLifeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
