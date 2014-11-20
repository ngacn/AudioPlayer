package client

import (
	"fmt"
	"net"
	"net/http"
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

func (cli *WarptenCli) getMethod(args ...string) (func(...string) error, bool) {
	camelArgs := make([]string, len(args))
	for i, s := range args {
		if len(s) == 0 {
			return nil, false
		}
		camelArgs[i] = strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	}
	methodName := "Cmd" + strings.Join(camelArgs, "")
	method := reflect.ValueOf(cli).MethodByName(methodName)
	if !method.IsValid() {
		return nil, false
	}
	return method.Interface().(func(...string) error), true
}

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

func NewWarptenCli(proto, addr string) *WarptenCli {
	tr := &http.Transport{}
	timeout := 32 * time.Second
	if proto == "unix" {
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
