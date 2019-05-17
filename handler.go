package main

import "fmt"

func handler(argSlice []string) {
	switch argSlice[0] {
	case "account":
		fmt.Println("account")
		accountHandler(argSlice[1:])
	case "tweet":
		fmt.Println("tweet")
	case "reset":
		fmt.Println("reset")
	case "help":
		fmt.Println("help")
	default:
		fmt.Println("invalid arguments.")
	}
}
