package main

import (
	"bufio"
	"fmt"
	"net"
	//"strconv"
)

func main (){
	r, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1313")
	conn, _ := net.DialTCP("tcp", nil, r)
	writer := bufio.NewWriter(conn)
    reader := bufio.NewReader(conn)
	for i:=0; i< 10; i++ {
		req := "Mensagem\n"
		_ , _ = writer.WriteString(req)
		writer.Flush()
		res, _ := reader.ReadString('\n')
		fmt.Println(res)
		
	}
}