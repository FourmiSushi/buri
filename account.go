package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/mrjones/oauth"
)

type keys struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type account struct {
	Screenname   string `json:"screen_name"`
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

func accountHandler(argSlice []string) {
	if len(argSlice) == 0 {
		argSlice = append(argSlice, "help")
	}
	switch argSlice[0] {
	case "add":
		addAccount()
	case "remove":
		removeAccount(argSlice[1:])
	case "list":
		listAccount()
	case "help":
		showAccountHelp()
	default:
		fmt.Println("invalid command.")
		showAccountHelp()
	}
}

func addAccount() {
	a := authorize()
	userSetting.Accounts = append(userSetting.Accounts, a)
	writeSettings(userSetting)
}

func removeAccount(argSlice []string) {
	removeNum, err := strconv.Atoi(argSlice[0])
	if err != nil {
		log.Fatal(err)
	}

	var newAccounts []account
	for i, s := range userSetting.Accounts {
		if i != removeNum {
			newAccounts = append(newAccounts, s)
		}
	}

	userSetting.Accounts = newAccounts
	userSetting.DefaultAccount = 0
	writeSettings(userSetting)

	fmt.Println("default account has been reset.")
	fmt.Printf("new default: %s\n", userSetting.Accounts[0].Screenname)
}

func listAccount() {
	for i, s := range userSetting.Accounts {
		if i != userSetting.DefaultAccount {
			fmt.Printf("%d    @%s\n", i, s.Screenname)
		} else {
			fmt.Printf("%d   *@%s\n", i, s.Screenname)
		}
	}
}

func showAccountHelp() {
	fmt.Print(
		`account is account maintenance command.

Usage:

	buri account <command>

The commands are:

	add                         add new account.
	remove <accountNumber>      remove an account from buri. 
	list                        show all accounts.
	help                        show this help.	
`)
	fmt.Println("")
}

func getKeys() (k keys) {
	bytes, err := ioutil.ReadFile("keys.json")
	if err != nil {
		log.Fatal(err)
	}
	var _k keys
	if err := json.Unmarshal(bytes, &_k); err != nil {
		log.Fatal(err)
	}

	return _k
}

func authorize() (a account) {
	k := getKeys()
	consumer := oauth.NewConsumer(
		k.ConsumerKey,
		k.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("open this url and enter the PIN code.")
	fmt.Println(url)
	fmt.Print("PIN > ")

	pin := ""
	fmt.Scanln(&pin)

	accessToken, err := consumer.AuthorizeToken(requestToken, pin)
	if err != nil {
		fmt.Println("Authenticate failed.")
		log.Fatal(err)
	}

	a.Screenname = accessToken.AdditionalData["screen_name"]
	a.AccessToken = accessToken.Token
	a.AccessSecret = accessToken.Secret

	return a
}
