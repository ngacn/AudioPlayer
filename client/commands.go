package client

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func (cli *WarptenCli) CmdHelp(args ...string) error {
	fmt.Fprintf(os.Stdout, "Usage: warpten [OPTIONS] COMMAND\n")
	flag.PrintDefaults()
	return nil
}

func (cli *WarptenCli) CmdVersion(args ...string) error {
	body, _, err := readBody(cli.call("GET", "/version", nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Warpten version: %s\n", body)
	return nil
}

func (cli *WarptenCli) CmdPlaylists(args ...string) error {
	body, _, err := readBody(cli.call("GET", "/playlists", nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Warpten playlists: %s\n", body)
	return nil

}

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
		body, _, err := readBody(cli.call("POST", "/playlist/new?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Create playlists %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	if *del {
		body, _, err := readBody(cli.call("DELETE", "/playlist/del?"+v.Encode(), nil))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Delete playlists %v: %s\n", cmd.Arg(0), body)
		return nil
	}

	body, _, err := readBody(cli.call("GET", "/playlist?"+v.Encode(), nil))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "Get playlist %v: %s\n", cmd.Arg(0), body)
	return nil
}
