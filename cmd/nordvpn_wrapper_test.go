package cmd_test

import (
	"bytes"
	"testing"
	"wvpn/cmd"

	"github.com/stretchr/testify/assert"
)

func TestNordVPNWrapper(t *testing.T) {
	vpn := cmd.NewNordVPN()
	vpn.Command("d")

	t.Run("toggle connection", func(t *testing.T) { ToggleCommand(t, "toggle", "t") })
	t.Run("help wrapped commands", HelpCommand)
}

func ToggleCommand(t *testing.T, toggleArgs ...string) {
	t.Helper()
	for _, arg := range toggleArgs {
		t.Run(arg, func(t *testing.T) {
			t.Run("connect", func(t *testing.T) {
				stdout := &bytes.Buffer{}
				cmd.NordVPNWrapper(stdout, arg)
				got := stdout.String()
				assert.Contains(t, got, "Connecting to")
				assert.Contains(t, got, "You are connected")
			})
			t.Run("disconnect", func(t *testing.T) {
				stdout := &bytes.Buffer{}
				cmd.NordVPNWrapper(stdout, arg)
				got := stdout.String()
				assert.Contains(t, got, "You are disconnected")
			})
		})
	}
}

func HelpCommand(t *testing.T) {
	t.Helper()
	for _, arg := range []string{"help", "h"} {
		t.Run(arg, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			cmd.NordVPNWrapper(stdout, arg)
			assert.Contains(t, stdout.String(), "Wrapped Commands")
		})
	}
}
