package main

import (
	"flag"
	"os"
	"wvpn/cmd"
)

func main() {
	flag.Parse()
	cmd.NordVPNWrapper(os.Stdout, flag.Args()...)
}
