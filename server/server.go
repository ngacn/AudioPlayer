package server

import (
	"bufio"
	"fmt"
	"net"
)

type Server interface {
	Serve()
	Close() error
}

type HttpServer struct {
	l net.Listener
}

func (s *HttpServer) Serve() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *HttpServer) Close() error {
	return s.l.Close()
}

func (s *HttpServer) handleConnection(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			break
		}
        if cmd == "EXIT\n" {
            fmt.Println("Exit ...")
            break
        } else {
            fmt.Printf(cmd)
        }
	}
	conn.Close()
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
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &HttpServer{l}, nil
}

func setupUnixHttp(addr string) (*HttpServer, error) {
	return nil, nil
}
