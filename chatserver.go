// a chat server implemented using net package

package main

import (
	"fmt"
	"net"
)

type client struct {
	id   string
	conn net.Conn
}

var allclients = make(map[string]client)
var uniqueId = 0

func generateId() string {
	uniqueId++
	return fmt.Sprintf("%d", uniqueId)
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		id := generateId()
		var newclient = client{id: id, conn: conn}
		fmt.Println("reading...")
		allclients[id] = newclient
		go handleConn(newclient)

	}
}

func handleConn(cli client) {
	b := make([]byte, 1024)
	for {
		cli.conn.Read(b)
		message := fmt.Sprint(cli.id + " : " + string(b))
		fmt.Println(message)
		for _, client := range allclients {
			if cli.conn == client.conn{
				continue
			}
			client.conn.Write([]byte(message))
		}
	}
	cli.conn.Close()
}
