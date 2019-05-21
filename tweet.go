package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func tweetHandler(argSlice []string) {
	if len(argSlice) == 0 {
		showTweetHelp()
	} else {
		i, err := strconv.Atoi(argSlice[0])
		if err != nil {
			log.Fatal(err)
		}
		tweetMode(i)
	}
}

func tweetMode(accountNum int) {
	fmt.Println("tweet mode")
	toEnd := false
	for {
		fmt.Print(" > ")

		body := ""
		fmt.Scanln(&body)

		if body == "" && toEnd {
			break
		}
		if body == "" {
			toEnd = true
		} else {
			tweet(accountNum, body)
		}
	}
}

func tweet(accountNum int, body string) {
	fmt.Println("tweet")
	k := getKeys()

	consumer := oauth1.NewConfig(k.ConsumerKey, k.ConsumerSecret)
	atoken := oauth1.NewToken(userSetting.Accounts[userSetting.DefaultAccount].AccessToken, userSetting.Accounts[userSetting.DefaultAccount].AccessSecret)
	httpClient := consumer.Client(oauth1.NoContext, atoken)
	client := twitter.NewClient(httpClient)
	_, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
}

func showTweetHelp() {
	fmt.Println("tweet help")
}
