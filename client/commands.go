package client

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

// 输出Usage
func (cli *WarptenCli) CmdHelp(args ...string) error {
	fmt.Fprintf(os.Stdout, "Usage: warpten [OPTIONS] COMMAND\n")
	flag.PrintDefaults()
	return nil
}

// 请求并输出版本号
func (cli *WarptenCli) CmdVersion(args ...string) error {
	body, _, err := readBody(cli.call("GET", "/version", nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Warpten version: %s\n", body)
	return nil
}

// 请求并输出所有播放列表
func (cli *WarptenCli) CmdPlaylists(args ...string) error {
	body, _, err := readBody(cli.call("GET", "/playlists", nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Warpten playlists: %s\n", body)
	return nil
}

// 请求并输出指定名字的播放列表， -a为增加新列表， -d为删除列表
func (cli *WarptenCli) CmdPlaylist(args ...string) error {
	cmd := cli.Subcmd("playlist", "[OPTIONS] ", "NAME", "Get playlist by name")
	add := cmd.Bool("a", false, "Create new playlist")
	del := cmd.Bool("d", false, "Delete playlist")

	if err := cmd.Parse(args); err != nil {
		return err
	}

	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}
	v := url.Values{}
	v.Set("name", cmd.Arg(0))

	// TODO: need refactor
	if *add {
		body, _, err := readBody(cli.call("POST", "/playlist/add?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Create playlist %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	if *del {
		body, _, err := readBody(cli.call("DELETE", "/playlist/del?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Delete playlist %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	body, _, err := readBody(cli.call("GET", "/playlist?"+v.Encode(), nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Get playlist %v: %s\n", cmd.Arg(0), body)
	return nil
}

// 请求并输出所有track
func (cli *WarptenCli) CmdTracks(args ...string) error {
	body, _, err := readBody(cli.call("GET", "/tracks", nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Warpten tracks: %s\n", body)
	return nil
}

// 请求并输出指定uuid的track信息， -a为增加新track， -d为删除某uuid的track
// 这两个操作都需要指定哪个播放列表中的track
// track并不储存在播放列表中， 只是方便移除相应的tag
func (cli *WarptenCli) CmdTrack(args ...string) error {
	cmd := cli.Subcmd("track", "[OPTIONS] ", "UUID", "Get track by uuid")
	pl := cmd.String("pl", "Default", "Add/Del track to/from playlist")
	add := cmd.Bool("a", false, "Create new track")
	del := cmd.Bool("d", false, "Delete track")

	if err := cmd.Parse(args); err != nil {
		return err
	}

	if cmd.NArg() < 1 {
		cmd.Usage()
		return nil
	}
	v := url.Values{}

	// TODO: need refactor
	if *add {
		v.Set("path", cmd.Arg(0))
		v.Set("playlist", *pl)
		body, _, err := readBody(cli.call("POST", "/track/add?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Create track %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	if *del {
		v.Set("uuid", cmd.Arg(0))
		v.Set("playlist", *pl)
		body, _, err := readBody(cli.call("DELETE", "/track/del?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Delete track %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	v.Set("uuid", cmd.Arg(0))
	body, _, err := readBody(cli.call("GET", "/track?"+v.Encode(), nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Get track %v: %s\n", cmd.Arg(0), body)
	return nil
}
