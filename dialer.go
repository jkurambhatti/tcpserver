package main

import (
	"io/ioutil"
	"net"
)

func main() {
	// conn, _ := net.Dial("tcp", "https://www.google.com")
	conn, _ := net.Dial("tcp", "localhost:8080")
	conn.Write([]byte("hello jayesh"))
	println(string(bs))
	conn.Close()
}
