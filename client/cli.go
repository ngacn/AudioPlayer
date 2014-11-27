// restful api的客户端程序
package client

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

type WarptenCli struct {
	proto     string
	addr      string
	transport *http.Transport
}

type StatusError struct {
	Status     string
	StatusCode int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("Status: %s, Code: %d", e.Status, e.StatusCode)
}

// 使用反射拼接命令行参数成具体的服务器端的函数
func (cli *WarptenCli) getMethod(args ...string) (func(...string) error, bool) {
	camelArgs := make([]string, len(args))
	for i, s := range args {
		if len(s) == 0 {
			return nil, false
		}
		camelArgs[i] = strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	}
	methodName := "Cmd" + strings.Join(camelArgs, "")
	// 查询WarptenCli下的方法
	method := reflect.ValueOf(cli).MethodByName(methodName)
	if !method.IsValid() {
		return nil, false
	}
	return method.Interface().(func(...string) error), true
}

// 解析命令行参数，并执行对应的命令
func (cli *WarptenCli) Cmd(args ...string) error {
	if len(args) > 1 {
		method, exists := cli.getMethod(args[:2]...)
		if exists {
			return method(args[2:]...)
		}
	}
	if len(args) > 0 {
		method, exists := cli.getMethod(args[0])
		if !exists {
			fmt.Println("WarptenCli: Command not found:", args[0])
			return cli.CmdHelp()
		}
		return method(args[1:]...)
	}
	return cli.CmdHelp()
}

// 处理二级命令
func (cli *WarptenCli) Subcmd(name, options, signature, description string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: warpten %s %s%s\n\n%s\n\n", name, options, signature, description)
		flags.PrintDefaults()
		os.Exit(2)
	}
	return flags
}

func NewWarptenCli(proto, addr string) *WarptenCli {
	tr := &http.Transport{}
	timeout := 32 * time.Second
	if proto == "unix" {
		// 本地unix domain socket通讯时不压缩数据，好像本地tcp时也不需要
		tr.DisableCompression = true
		tr.Dial = func(_, _ string) (net.Conn, error) {
			return net.DialTimeout(proto, addr, timeout)
		}
	} else {
		tr.Dial = (&net.Dialer{Timeout: timeout}).Dial
	}

	return &WarptenCli{
		proto:     proto,
		addr:      addr,
		transport: tr,
	}
}
