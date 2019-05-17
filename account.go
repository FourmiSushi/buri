package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mrjones/oauth"
)

type keys struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type access struct {
	Screenname   string `json:"screen_name"`
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

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
	_a := authorize()
	b, err := json.Marshal(_a)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("~/.buritweet", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, string(b))
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

func authorize() (a access) {
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
	fmt.Println("open this url and enter the PIN code.")
	fmt.Println(url)
	fmt.Print("PIN >")
	pin := ""
	fmt.Scanln(&pin)
	accessToken, err := consumer.AuthorizeToken(requestToken, pin)
	if err != nil {
		fmt.Println("(Authenticate faild.)")
		log.Fatal(err)
	}

	a.Screenname = accessToken.AdditionalData["screen_name"]
	a.AccessToken = accessToken.Token
	a.AccessSecret = accessToken.Secret

	return a
}
