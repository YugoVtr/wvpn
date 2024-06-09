package cmd

import (
	"fmt"
	"io"
	"wvpn/net"
	"wvpn/os"
	"wvpn/vpn"
)

// NordVPNWrapper encapsulates NordVPN.
// It interprets the command line arguments and calls the appropriate NordVPN functions.
// Any unrecognized command is passed directly to NordVPN.
// Commands:
//
//	toggle, t				- Toggle the connection on/off. If off connect with random server.
//	ip, i						- Get the public IP address of the user.
//	mesh, m					- Toggle meshnet connection on/off.
//	peer, p					- Allow or deny all roles to a peer.
func NordVPNWrapper(w io.Writer, args ...string) {
	nordvpn := NewNordVPN()
	out, err := "", error(nil)
	switch Args(args).At(0) {
	case "toggle", "t":
		out, err = nordvpn.ToggleConnection()
	case "ip", "i":
		out = fmt.Sprintf("%sYour IP is %s", out, net.IP())
	case "mesh", "m":
		out, err = nordvpn.MeshNet().ToggleConnection()
	case "peer", "p":
		out, err = nordvpn.MeshNet().AllowAllToPeer(
			Args(args).At(1),
			Args(args).At(2) == "allow",
		)
	case "help", "h":
		out = nordvpn.Help()
	default:
		out, err = nordvpn.Command(args...)
	}
	os.Print(w, out, err)
}

// NewNordVPN creates a new NordVPN instance.
func NewNordVPN() vpn.NordVPN {
	return vpn.NewNordVPN(os.Command)
}

// Args store command line arguments.
type Args []string

// At returns the element at position p of the Args slice.
// If args have no element, it returns an empty string.
func (args Args) At(p int) string {
	if p < len(args) {
		return args[p]
	}
	return ""
}
