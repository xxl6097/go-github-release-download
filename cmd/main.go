package main

import (
	"github.com/xxl6097/go-github-publish-release/github"
	"log"
)

func main() {
	path, err := github.Download("./", "AuGoService", "0.0.0", "xxl6097", "go-service-framework")
	if err == nil {
		log.Println(path)
	}
}
