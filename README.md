# WVPN - A NordVPN Wrapper CLI

**wvpn** encapsulates NordVPN. It interprets the command line arguments and calls the appropriate NordVPN functions. Any unrecognized command is passed directly to NordVPN.

## Requirements
- [NordVPN](https://nordvpn.com/download/linux)

## Installation
```bash
  $ go install github.com/yugovtr/wvpn@latest
```

## Usage
```bash
  $ wvpn [command] [arguments]
```
Example:
```bash
  $ wvpn help

  Usage: wvpn [global options] command [command options] [arguments...]

  Commands:
      account               Shows account information
      cities                Shows a list of cities where servers are available
      connect, c            Connects you to VPN
      countries             Shows a list of countries where servers are available
      disconnect, d         Disconnects you from VPN
      groups                Shows a list of available server groups
      login                 Logs you in
      logout                Logs you out
      rate                  Rates your last connection quality (1-5)
      register              Registers a new user account
      set, s                Sets a configuration option
      settings              Shows current settings
      status                Shows connection status
      version               Shows the app version
      allowlist, whitelist  Adds or removes an option from the allowlist
      meshnet, mesh         Meshnet is a way to safely access other devices, no matter where in the world they are. Once set up, Meshnet functions just like a secure local area network (LAN) — it connects devices directly. It also allows securely sending files to other devices. Use the "nordvpn set meshnet on" command to enable Meshnet. Learn more: https://meshnet.nordvpn.com/
      fileshare             Transfer files of any size between Meshnet peers securely and privately
      help, h               Shows a list of commands or help for one command

  Wrapped Commands:
      toggle, t           Connects to a random country if not connected, otherwise disconnects.
      ip, i               Shows public IP address of the user.
      mesh, m             Connects to meshnet if not connected, otherwise disconnects.
      peer, p             Allow or deny all roles to a peer.

  Global options:
    --help, -h     Show help
    --version, -v  Print the version
```
