// 程序的入口文件，目前需要-d后台启动一次，开启服务器，然后不带-d执行向服务器发送命令,
// 实际的做法应该是执行一次服务器启动为守护程序，然后fork出gui进程。
package main

import (
	"flag"
	"fmt"
	"strings"
	"warpten/client"
	"warpten/player"
	"warpten/server"
)

func main() {
	const (
		DEFAULTHTTPHOST   = "127.0.0.1:7478"
		DEFAULTUNIXSOCKET = "/tmp/warpten.sock"
	)

	// 设置命令行参数
	var (
		flDaemon = flag.Bool("d", false, "Enable daemon mode")
		flTcp    = flag.Bool("t", false, "Enable the TCP socket")
	)

	// 解析命令行参数
	flag.Parse()

	// 默认使用unix domin socket连接，带-t时用tcp
	defaultHost := fmt.Sprintf("unix://%s", DEFAULTUNIXSOCKET)
	if *flTcp {
		defaultHost = fmt.Sprintf("tcp://%s", DEFAULTHTTPHOST)
	}

	protoAddrParts := strings.SplitN(defaultHost, "://", 2)

	// 说明是启动服务器端的功能
	if *flDaemon {
		// 初始化播放器
		player.Init()
		// 创建restful api服务器端
		srv, err := server.NewWarptenSrv(protoAddrParts[0], protoAddrParts[1])
		if err != nil {
			fmt.Println("WarptenSrv: Couldn't create WarptenSrv")
			return
		}
		// 开始监听并处理请求
		srv.Serve()
		// 程序退出
		return
	}

	// 创建restful api客户端程序
	var cli *client.WarptenCli = client.NewWarptenCli(protoAddrParts[0], protoAddrParts[1])
	// 解析命令，并向服务器发送请求
	if err := cli.Cmd(flag.Args()...); err != nil {
		fmt.Println("WarptenCli: =>", err)
	}
}
