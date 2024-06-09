package cmd

import (
	"fmt"
	"wvpn/net"
	"wvpn/vpn"
)

type Strategy func(nordvpn vpn.NordVPN, args ...string) (string, error)

func Toggle(nordvpn vpn.NordVPN, _ ...string) (string, error) {
	return nordvpn.ToggleConnection()
}

func IP(nordvpn vpn.NordVPN, _ ...string) (string, error) {
	return fmt.Sprintf("Your IP is %s", net.IP()), nil
}

func Mesh(nordvpn vpn.NordVPN, _ ...string) (string, error) {
	return nordvpn.MeshNet().ToggleConnection()
}

func Peer(nordvpn vpn.NordVPN, args ...string) (string, error) {
	return nordvpn.MeshNet().AllowAllToPeer(
		Args(args).At(1),
		Args(args).At(2) == "allow",
	)
}

func Help(nordvpn vpn.NordVPN, _ ...string) (string, error) {
	return nordvpn.Help(), nil
}

var commands = map[string]Strategy{
	"toggle": Toggle,
	"t":      Toggle,
	"ip":     IP,
	"i":      IP,
	"mesh":   Mesh,
	"m":      Mesh,
	"peer":   Peer,
	"p":      Peer,
	"help":   Help,
	"h":      Help,
}
