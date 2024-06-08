package main

import (
	"fmt"
	"wvpn/net"
	"wvpn/os"
	"wvpn/vpn"
)

func main() {
	nordvpn := vpn.NewNordVPN(os.Command)
	os.Print(nordvpn.ToggleConnection())
	fmt.Printf("Your IP is %s", net.IP())
}
