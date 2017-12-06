package main

import (
	"fmt"
	"net"
	"os"
)

func recvConnMsg(conn net.Conn)  {
	buf:=make([]byte,1024)
	defer conn.Close()

	for{
		n,err := conn.Read(buf)

		if err != nil{
			fmt.Println("conn closed")
			return
		}

		fmt.Println("recv msg:", string(buf[0:n]))
	}


}

func main() {

	fmt.Println("running.....tcpserver")

	listen_sock,err := net.Listen("tcp","localhost:10000")
	if  err != nil {
		fmt.Println("Error: %s", err.Error())

		os.Exit(1)
	}

	defer listen_sock.Close()

	for{
		new_conn, err :=listen_sock.Accept()

		if err != nil{
			continue
		}

		go recvConnMsg(new_conn)


	}









}
