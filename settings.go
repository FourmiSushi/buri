package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type setting struct {
	Accounts       []account `json:"accounts"`
	DefaultAccount int       `json:"default"`
}

func getSettingsPath() (path string) {
	userDir, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path = userDir.HomeDir + "/.buritweet"
	return path
}

func isSettingsExist() (isExist bool) {
	_, err := os.Stat(getSettingsPath())
	return err == nil
}

func createSettings() {
	file, err := os.Create(getSettingsPath())
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func writeSettings(s setting) {
	file, err := os.OpenFile(getSettingsPath(), os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(file, string(_b))
}

func readSettings() (s setting) {
	bytes, err := ioutil.ReadFile(getSettingsPath())
	if err != nil {
		log.Fatal(err)
	}
	var _s setting
	if err := json.Unmarshal(bytes, &_s); err != nil {
		log.Fatal(err)
	}
	return _s
}

func resetSettings(){
	file, err := os.OpenFile(getSettingsPath(), os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprint(file, "")
}
