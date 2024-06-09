package vpn

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// NewNordVPN creates a new NordVPN instance.
func NewNordVPN(c Commander) NordVPN {
	return NordVPN{cmd: c, name: "nordvpn"}
}

// Commander is a function that executes a command and returns the stdout and error.
type Commander func(string, ...string) (string, error)

type NordVPN struct {
	cmd  Commander
	name string
}

// Command executes a NordVPN command and returns output and error.
// Needs nordvpn in system PATH.
func (vpn NordVPN) Command(args ...string) (string, error) {
	return vpn.cmd(vpn.name, args...)
}

// Connected returns true if the VPN is connected.
func (vpn NordVPN) Connected() bool {
	status, _ := vpn.Command("status")
	return strings.Contains(status, "Connected")
}

// Countries returns a list of countries available for connection.
func (vpn NordVPN) Countries() []string {
	countries, _ := vpn.Command("countries")
	countriesList := strings.Split(strings.TrimSpace(countries), ",")
	return countriesList
}

// Connect NordVPN wrapper for the connect command.
func (vpn NordVPN) Connect(country string) (string, error) {
	return vpn.Command("c", country)
}

// Disconnect NordVPN wrapper for the disconnect command.
func (vpn NordVPN) Disconnect() (string, error) {
	return vpn.Command("d")
}

// RandomCountry returns a random country from the list of countries.
func (vpn NordVPN) RandomCountry() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	countries := vpn.Countries()
	country := countries[r.Intn(len(countries))]
	return strings.TrimSpace(country)
}

// ToggleConnection connects to a random country if not connected, otherwise disconnects.
func (vpn NordVPN) ToggleConnection() (string, error) {
	if !vpn.Connected() {
		country := vpn.RandomCountry()
		return vpn.Connect(country)
	}
	return vpn.Disconnect()
}

// MeshNet returns a MeshNet instance.
func (vpn NordVPN) MeshNet() MeshNet {
	return NewMeshNet(vpn)
}

// Help add custom commands to the NordVPN help.
func (vpn NordVPN) Help() string {
	const repacleAt = "Global options:"
	nordVPNHelp, _ := vpn.Command("h")
	wVPNHelp := fmt.Sprintf("%s\n\n%s", vpn.help(), repacleAt)
	nordVPNHelp = strings.ReplaceAll(nordVPNHelp, "Usage: nordvpn", "Usage: wvpn")
	nordVPNHelp = strings.ReplaceAll(nordVPNHelp, repacleAt, wVPNHelp)

	lines := strings.Split(nordVPNHelp, "\n")
	return strings.Join(lines[3:len(lines)-4], "\n")
}

func (vpn NordVPN) help() string {
	return "\033[33mWrapped Commandss\033[0m:" + `
     toggle, t           Connects to a random country if not connected, otherwise disconnects.
     ip, i               Shows public IP address of the user.
     mesh, m             Connects to meshnet if not connected, otherwise disconnects.
     peer, p             Allow or deny all roles to a peer.`
}
