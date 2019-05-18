package main

import (
	"os"
)

var userSetting setting

func main() {
	args := os.Args[1:]

	if isSettingsExist() {
		userSetting = readSettings()
	} else {
		createSettings()
	}

	if len(args) == 0 {
		args = append(args, "help")
	}

	handler(args)
}
