//go:build linux || darwin
// +build linux darwin

package notification

import "github.com/ctcpip/notifize"

func Push(msg string) {
	notifize.Display("Albionshit", msg, false, "")
}
