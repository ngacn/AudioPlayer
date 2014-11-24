package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"syscall"
	"warpten/player"
)

type HttpApiFunc func(w http.ResponseWriter, r *http.Request) error

func createRouter() (*http.ServeMux, error) {
	r := http.NewServeMux()
	m := map[string]map[string]HttpApiFunc{
		"GET": {
			"/version":   getVersion,
			"/playlists": getPlaylists,
			"/playlist":  getPlaylist,
		},
		"POST": {
			"/playlist/new": newPlaylist,
		},
		"DELETE": {
			"/playlist/del": delPlaylist,
		},
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
	fmt.Fprintf(w, player.Version())
	return nil
}

func getPlaylists(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	pls := player.Playlists()
	b, err := json.Marshal(pls)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}

func getPlaylist(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	name := r.Form.Get("name")
	if pl, exists := player.Playlist(name); exists {
		b, err := json.Marshal(pl)
		if err != nil {
			return err
		}
		w.Write(b)
		return nil
	}
	fmt.Fprintf(w, "nil")
	return nil
}

func newPlaylist(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	name := r.Form.Get("name")
	if err := player.NewPlaylist(name); err != nil {
		fmt.Fprintf(w, name+" exists")
		return nil
	}
	fmt.Fprintf(w, "success")
	return nil
}

func delPlaylist(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	name := r.Form.Get("name")
	if err := player.DelPlaylist(name); err != nil {
		fmt.Fprintf(w, name+" not exists")
		return nil
	}
	fmt.Fprintf(w, "success")
	return nil
}

func makeHttpHandler(localMethod string, localRoute string, handlerFunc HttpApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r); err != nil {
			fmt.Println("WarptenSrv: Handler for %s %s returned error: %s", localMethod, localRoute, err)
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

func parseForm(r *http.Request) error {
	if r == nil {
		return nil
	}
	if err := r.ParseForm(); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}
