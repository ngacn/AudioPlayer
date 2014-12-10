// 程序的入口文件，-d可以启动服务器守护程序
package main

import (
	"flag"
	"fmt"
	"strings"
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
}
