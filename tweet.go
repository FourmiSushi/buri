package main

import (
	"fmt"
	"log"
	"strconv"
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
}

func tweet(accountNum int) {
	fmt.Println("tweet")
}

func showTweetHelp() {
	fmt.Println("tweet help")
}
