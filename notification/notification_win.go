//go:build windows
// +build windows

package notification

import (
	"github.com/ao-data/albiondata-client/log"
	toast "gopkg.in/toast.v1"
)

func Push(msg string) {
	note := toast.Notification{AppID: "Albionshit", Title: "Albionshit", Message: msg}

	err := note.Push()
	if err != nil {
		log.Error(err)
	}
}
