package main

import "fmt"

func handler(argSlice []string) {
	switch argSlice[0] {
	case "account":
		fmt.Println("account")
	case "tweet":
		fmt.Println("tweet")
	case "reset":
		fmt.Println("reset")
	}
}
