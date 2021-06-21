package main

import (
	"io/ioutil"
	"os"

	"github.com/Yeuoly/kirakira/core"
	"github.com/Yeuoly/kirakira/github"
)

func GetProxy() string {
	proxy, err := ioutil.ReadFile("proxy")
	if err != nil {
		return ""
	}
	return string(proxy)
}

func GetAccessKey() string {
	key, err := ioutil.ReadFile("key")
	if err != nil {
		return ""
	}
	return string(key)
}

func main() {
	client := github.GithubClient{}
	client.Init(GetProxy(), GetAccessKey())

	str, err := core.ReplaceMD(&client, "README_template.md")

	if err != nil {
		panic(err)
	}

	os.WriteFile("README.md", []byte(str), 0644)
}
