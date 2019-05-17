package main

import "fmt"

func accountHandler(argSlice []string) {
	fmt.Println(argSlice)
	switch argSlice[0] {
	case "add":
		fmt.Println("add")
		addAccount()
	case "remove":
		fmt.Println("remove")
		removeAccount(argSlice[1:])
	case "list":
		fmt.Println("list")
		listAccount()
	case "help":
		fmt.Println("help")
		helpAccount()
	default:
		fmt.Println("invalid arguments")
		helpAccount()
	}
}

func addAccount() {
	fmt.Println("addAccount")
}

func removeAccount(argSlice []string) {
	fmt.Println("removeAccount")
	fmt.Println(argSlice)
}

func listAccount() {
	fmt.Println("listAccount")
}

func helpAccount() {
	fmt.Println("helpAccount")
}
