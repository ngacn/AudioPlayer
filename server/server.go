// 这是restful api的服务器端程序
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

// 这相当于一个声明一个回调函数， 不同的请求用不同的函数处理
type HttpApiFunc func(w http.ResponseWriter, r *http.Request) error

// 创建不同的请求和回调函数的对应关系
func createRouter() (*http.ServeMux, error) {
	r := http.NewServeMux()
	m := map[string]map[string]HttpApiFunc{
		"GET": {
			"/version":   getVersion,
			"/playlists": getPlaylists,
			"/playlist":  getPlaylist,
			"/tracks":    getTracks,
			"/track":     getTrack,
		},
		"POST": {
			"/playlist/add": addPlaylist,
			"/track/add":    addTrack,
		},
		"DELETE": {
			"/playlist/del": delPlaylist,
			"/track/del":    delTrack,
		},
	}

	for method, routes := range m {
		for route, fct := range routes {
			localRoute := route
			localFct := fct
			localMethod := method
			// 给回调函数包装了一层
			f := makeHttpHandler(localMethod, localRoute, localFct)
			// 绑定请求和回调函数
			r.HandleFunc(localRoute, f)
		}
	}

	return r, nil
}

// 获取播放器版本
func getVersion(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, player.Version())
	return nil
}

// 获取所有播放列表
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

// 获取指定名字的播放列表
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

// 添加播放列表
func addPlaylist(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	name := r.Form.Get("name")
	if err := player.AddPlaylist(name); err != nil {
		fmt.Fprintf(w, name+" exists")
		return nil
	}
	fmt.Fprintf(w, "success")
	return nil
}

// 删除播放列表
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

// 获取所有tracks
func getTracks(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	tks := player.Tracks()
	b, err := json.Marshal(tks)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}

// 获取指定uuid的track
func getTrack(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	uuid := r.Form.Get("uuid")
	if tk, exists := player.Track(uuid); exists {
		b, err := json.Marshal(tk)
		if err != nil {
			return err
		}
		w.Write(b)
		return nil
	}
	fmt.Fprintf(w, "nil")
	return nil
}

// 添加指定路径的track到某个播放列表， 比如在gui中拖拽一个文件到播放列表
func addTrack(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	path, playlist := r.Form.Get("path"), r.Form.Get("playlist")
	if err := player.AddTrack(path, playlist); err != nil {
		fmt.Fprintf(w, err.Error())
		return nil
	}
	fmt.Fprintf(w, "success")
	return nil
}

// 删除指定uuid的track
func delTrack(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if err := parseForm(r); err != nil {
		return err
	}
	uuid := r.Form.Get("uuid")
	if err := player.DelTrack(uuid); err != nil {
		fmt.Fprintf(w, err.Error())
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

// 服务器接口
type Server interface {
	Serve() error
	Close() error
}

// 服务器结构体
type HttpServer struct {
	srv *http.Server
	l   net.Listener
}

// 开始处理请求
func (s *HttpServer) Serve() error {
	return s.srv.Serve(s.l)
}

// 关闭
func (s *HttpServer) Close() error {
	return s.l.Close()
}

// 创建tcp或unix domain socket
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

// 解析request中的参数
func parseForm(r *http.Request) error {
	if r == nil {
		return nil
	}
	if err := r.ParseForm(); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return err
	}
	return nil
}
