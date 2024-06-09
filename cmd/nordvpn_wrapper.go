package cmd

import (
	"io"

	"github.com/yugovtr/wvpn/os"
	"github.com/yugovtr/wvpn/vpn"
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
	defer func() { os.Print(w, out, err) }()

	if caller, ok := commands[Args(args).At(0)]; ok {
		out, err = caller(nordvpn, args...)
		return
	}
	out, err = nordvpn.Command(args...)
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
