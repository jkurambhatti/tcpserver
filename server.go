package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		io.WriteString(conn, "hello")
		fmt.Printf("%v", conn.RemoteAddr())
		conn.Close()
	}
}
