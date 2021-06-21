package main

import (
	"fmt"

	"github.com/Yeuoly/kirakira/github"
)

func main() {
	client := github.GithubClient{}
	client.Init("http://127.0.0.1:7890", "")

	exist := client.ProfileExists("BestLiy")

	if exist {
		fmt.Println("yes!")
	} else {
		fmt.Println("no")
	}
}
