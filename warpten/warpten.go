package main

import (
	"fmt"
	"warpten/server"
)

func main() {
	srv, err := server.NewWarptenSrv("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Couldn't create WarptenSrv")
		return
	}
	srv.Serve()
}
