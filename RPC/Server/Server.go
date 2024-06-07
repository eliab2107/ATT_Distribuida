package main

import (
	"bufio"
	"fmt"
	"net"
)

func conecta(server net.Conn){
	for {
		message , _ :=  bufio.NewReader(server).ReadString('\n')
		res := "funcionando\n"
		fmt.Println("deu bom ", message)
		_, _ = server.Write([]byte(res))
	}
	
}

func main(){
	server, _ := net.Listen("tcp", "127.0.0.1:1313")
	fmt.Println("Listen on IP and port: 127.0.0.1:1313")
	for{
		conn, _ := server.Accept()
		defer conn.Close()
		fmt.Println("conexao aceita")
		go conecta(conn)
	}
}