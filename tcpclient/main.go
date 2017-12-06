package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:10000")

	if err != nil {

		fmt.Println("Error ：",err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	conn.Write([]byte("hello world!"))
	fmt.Println("send msg...")



}
