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
	var user string
	var req string
	fmt.Scanln(&user)
	_ , _ = writer.WriteString(user+"\n")
	writer.Flush()
	for {
		fmt.Scanln(&req)
		_ , _ = writer.WriteString(req+"\n")
		writer.Flush()
		res, _ := reader.ReadString('\n')
		fmt.Println(res)
	}
}