package main

import (
	"bufio"
	"fmt"
	"net"
)

func conecta(server net.Conn, user string){
	var res string
	for {
		_, _ =  bufio.NewReader(server).ReadString('\n')
		res = "Confirmação de recebimento de " +  user
		_, _ = server.Write([]byte(res))
	}
	
}

func main(){
	r, _ :=net.ResolveTCPAddr("tcp","127.0.0.1:1313" )
	server, _ := net.ListenTCP("tcp", r)
	fmt.Println("Listen on IP and port: 127.0.0.1:1313")
	for{
		conn, _ := server.Accept()
		defer conn.Close()
		user , _ :=  bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Conexao iniciada com ", user)
		go conecta(conn, user)
	}
}