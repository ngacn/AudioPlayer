package main

import (
	"flag"
	"fmt"
	"strings"
	"warpten/client"
	"warpten/server"
)

func main() {
	const (
		DEFAULTHTTPHOST   = "127.0.0.1:7478"
		DEFAULTUNIXSOCKET = "/tmp/warpten.sock"
	)

	var (
		flDaemon = flag.Bool("d", false, "Enable daemon mode")
		flTcp    = flag.Bool("t", false, "Enable the TCP socket")
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

	var cli *client.WarptenCli = client.NewWarptenCli(protoAddrParts[0], protoAddrParts[1])
	if err := cli.Cmd(flag.Args()...); err != nil {
		fmt.Println("WarptenCli: =>", err)
	}
}
