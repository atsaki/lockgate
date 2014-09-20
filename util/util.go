package util

import (
	"os/user"
	"strings"
)

func ExpandPath(path string) string {
	usr, _ := user.Current()
	home := usr.HomeDir

	if strings.HasPrefix(path, "~/") {
		path = strings.Replace(path, "~/", home+"/", 1)
	}
	return path
}