package vpn

import (
	"strings"
)

// NewMeshNet creates a new MeshNet instance.
func NewMeshNet(nord NordVPN) MeshNet {
	return MeshNet{NordVPN: nord, name: "meshnet"}
}

// MeshNet is a subcommand of NordVPN provinder.
type MeshNet struct {
	NordVPN
	name string
}

// Command executes a meshnet command and returns output and error.
// Example: nordvpn meshnet set on
func (mesh MeshNet) Command(args ...string) (string, error) {
	args = append([]string{"meshnet"}, args...)
	return mesh.NordVPN.Command(args...)
}

// Connect to the meshnet.
func (mesh MeshNet) Connect() (string, error) {
	return mesh.NordVPN.Command("set", mesh.name, "on")
}

// Disconnect from the meshnet.
func (mesh MeshNet) Disconnect() (string, error) {
	return mesh.NordVPN.Command("set", mesh.name, "off")
}

// Connected returns true if the meshnet is connected.
// When nordvpn settings have meshnet status as enabled.
func (mesh MeshNet) Connected() bool {
	status, _ := mesh.NordVPN.Command("settings")
	return strings.Contains(status, "Meshnet: enabled")
}

// ToggleConnection connects to the meshnet if not connected, otherwise disconnects.
func (mesh MeshNet) ToggleConnection() (string, error) {
	if !mesh.Connected() {
		mesh.NordVPN.Command("c")
		return mesh.Connect()
	}
	mesh.NordVPN.Command("d")
	return mesh.Disconnect()
}

// AllowAllToPeer allows or denies all roles to a peer.
// Roles: incoming, routing, local, fileshare, auto-accept.
// Returns the list of peers with roles permissions.
func (mesh MeshNet) AllowAllToPeer(peer string, allow bool) (string, error) {
	permission, autoaccept, cmd := "deny", "disable", "peer"
	if allow {
		permission, autoaccept = "allow", "enable"
	}
	roles := []string{"incoming", "routing", "local", "fileshare"}
	for _, role := range roles {
		mesh.Command(cmd, role, permission, peer)
	}
	mesh.Command(cmd, "auto-accept", autoaccept, peer)
	return mesh.Command(cmd, "list")
}
