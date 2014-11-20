package client

import (
	"flag"
	"fmt"
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
