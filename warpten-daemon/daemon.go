// 程序的入口文件，-d可以启动服务器守护程序
package main

import (
	"flag"
	"fmt"
	"warpten/player"
	"warpten/server"
)

func main() {
	const (
		DEFAULTHTTPHOST = "127.0.0.1:7478"
	)

	// 设置命令行参数
	var (
		flDaemon = flag.Bool("d", false, "Enable daemon mode")
	)

	// 解析命令行参数
	flag.Parse()

	// 说明是启动服务器端的功能
	if *flDaemon {
		// 初始化播放器
		player.New()
		// 创建restful api服务器端
		srv, err := server.NewWarptenSrv("tcp", DEFAULTHTTPHOST)
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
