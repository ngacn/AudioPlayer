// 播放器的入口文件， 会fork出一个进程启动服务器守护程序
package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"log"
	"os"
	"os/exec"
)

func main() {
	// 启动服务器守护程序
	cmd := exec.Command("warpten-daemon", "-d")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 启动gui
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	// 关闭服务器守护程序
	cmd.Process.Signal(os.Interrupt)
}

func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}
