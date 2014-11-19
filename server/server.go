package server

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
)

type HttpApiFunc func(w http.ResponseWriter, r *http.Request) error

func createRouter() (*http.ServeMux, error) {
	r := http.NewServeMux()
	m := map[string]map[string]HttpApiFunc{
		"GET": {
			"/version": getVersion,
		},
		"POST":   {},
		"DELETE": {},
	}
	for method, routes := range m {
		for route, fct := range routes {
			localRoute := route
			localFct := fct
			localMethod := method
			f := makeHttpHandler(localMethod, localRoute, localFct)
			r.HandleFunc(localRoute, f)
		}
	}

	return r, nil
}

func getVersion(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "yoyoyoyo")
	return nil
}

func makeHttpHandler(localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			fmt.Println("WarptenSrv: handlerFunc failed.")
		}
	}
}

type Server interface {
	Serve() error
	Close() error
}

type HttpServer struct {
	srv *http.Server
	l   net.Listener
}

func (s *HttpServer) Serve() error {
	return s.srv.Serve(s.l)
}

func (s *HttpServer) Close() error {
	return s.l.Close()
}

func NewWarptenSrv(proto, addr string) (Server, error) {
	switch proto {
	case "tcp":
		return setupTcpHttp(addr)
	case "unix":
		return setupUnixHttp(addr)
	default:
		return nil, fmt.Errorf("Invalid protocol format.")
	}
}

func setupTcpHttp(addr string) (*HttpServer, error) {
	r, err := createRouter()
	if err != nil {
		return nil, err
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &HttpServer{&http.Server{Addr: addr, Handler: r}, l}, nil
}

func setupUnixHttp(addr string) (*HttpServer, error) {
	r, err := createRouter()
	if err != nil {
		return nil, err
	}

	if err := syscall.Unlink(addr); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	mask := syscall.Umask(0777)
	defer syscall.Umask(mask)

	l, err := net.Listen("unix", addr)
	if err != nil {
		return nil, err
	}
	if err := os.Chmod(addr, 0660); err != nil {
		return nil, err
	}

	return &HttpServer{&http.Server{Addr: addr, Handler: r}, l}, nil
}
