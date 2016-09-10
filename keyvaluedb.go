//  a simple REDIS like, in memory database which uses key : value pair to store data
//  uses only net package to handle tcp connections
//  note that both key and value is a string
//  supports multiple connections, all have access to the redisDB

//  the database uses the following commands
//  GET  "key" : to get the value of the key
//  SET  "key"  "value" : to set the value for any given key
//  DEL  "key" : to delete a given key

package main

import (
	"fmt"
	"net"
	"strings"
)

var redisDB = make(map[string]string)

func redisserver(c net.Conn) {
	var bs = make([]byte, 1024)
	for {
		n ,_ :=c.Read(bs)
		cmd := string(bs[:n]) // to exclude the newline character
		splitcmd := strings.Fields(cmd)  // splits on one or more consecutive spaces and newline characters
		fmt.Println(len(splitcmd))
		fmt.Println(splitcmd)
		switch splitcmd[0] {
		case "GET":
			if len(splitcmd) != 2{
				c.Write([]byte("invalid request format \n"))
				c.Write([]byte("usage : GET key\n"))
				break
			}
			key := splitcmd[1]
			if val,ok :=redisDB[key];ok{
				c.Write([]byte(val))
				c.Write([]byte("\n"))
			} else {
				c.Write([]byte("no record found\n"))
			}
		case "SET":
			if len(splitcmd) != 3{
				c.Write([]byte("invalid request format \n"))
				c.Write([]byte("usage : SET key value\n"))
				break
			}
			key := splitcmd[1]
			val := splitcmd[2]

			redisDB[key] = val
			c.Write([]byte("OK\n"))
		case "DEL":
			if len(splitcmd) != 2{
				c.Write([]byte("invalid request format \n"))
				c.Write([]byte("usage : DEL key\n"))
				break
			}
			key := splitcmd[1]
			if _,ok :=redisDB[key];ok{
				delete(redisDB,key)
				c.Write([]byte("OK\n"))
			} else {
				c.Write([]byte("no record found\n"))
			}
		default:
			c.Write([]byte("invalid request\n"))

		}
		fmt.Println(cmd)
	}
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		fmt.Println("got a new connection")
		if err != nil {
			fmt.Println(err)
			return
		}

		go redisserver(conn)
	}

}
