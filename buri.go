package main

import (
	"os"
)

var userSetting setting

func main() {
	args := os.Args[1:]

	if isSettingsExist() {
		userSetting = readSettings()
	}

	handler(args)
}
