package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "send-cmd", "PLAY", "a command")
	flag.Parse()
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Couldn't connect to server")
		return
	}
	fmt.Fprintln(conn, cmd)
}
