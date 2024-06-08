package net

import (
	"io"
	"net/http"
)

const CheckIPURL = "http://checkip.amazonaws.com"

// IP returns the public IP address of the user.
func IP() string {
	r, err := http.Get(CheckIPURL)
	if err != nil {
		return ""
	}
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	return string(body)
}
