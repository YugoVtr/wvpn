package main

import (
	"flag"
	"os"

	"github.com/yugovtr/wvpn/cmd"
)

func main() {
	flag.Parse()
	cmd.NordVPNWrapper(os.Stdout, flag.Args()...)
}
