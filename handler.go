package main

import "fmt"

func handler(argSlice []string) {
	switch argSlice[0] {
	case "account":
		accountHandler(argSlice[1:])
	case "tweet":
		tweetHandler(argSlice[1:])
	case "reset":
		resetSettings()
	case "help":
		showHelp()
	default:
		fmt.Print("invalid command or arguments.\n\n")
		showHelp()
	}
}

func showHelp() {
	fmt.Print(
		`buri is Twitter Client for Shell.

Usage:

	buri <command> [arguments]

The commands are:

	account     account maintenance.
	tweet       start tweet mode.
	reset       delete all settings.
	help        show this help.
	
Use "buri <command> help" for more information about a command.
`)
	fmt.Println("")
}
