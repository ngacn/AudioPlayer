package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"warpten/server"
)

func main() {
	const (
		DEFAULTHTTPHOST   = "127.0.0.1:7478"
		DEFAULTUNIXSOCKET = "/var/run/warpten.sock"
	)

	var (
		flDaemon = flag.Bool("d", false, "Enable daemon mode")
		flTcp    = flag.Bool("t", false, "Enable the TCP socket")
		flCmd    = flag.String("s", "PLAY", "Send a command to server")
	)

	flag.Parse()

	defaultHost := fmt.Sprintf("unix://%s", DEFAULTUNIXSOCKET)
	if *flTcp {
		defaultHost = fmt.Sprintf("tcp://%s", DEFAULTHTTPHOST)
	}

	protoAddrParts := strings.SplitN(defaultHost, "://", 2)

	if *flDaemon {
		srv, err := server.NewWarptenSrv(protoAddrParts[0], protoAddrParts[1])
		if err != nil {
			fmt.Println("WarptenSrv: Couldn't create WarptenSrv")
			return
		}
		srv.Serve()
		return
	}

	conn, err := net.Dial(protoAddrParts[0], protoAddrParts[1])
	if err != nil {
		fmt.Println("WarptenCli: Couldn't connect to server")
		return
	}
	fmt.Fprintln(conn, *flCmd)
}
