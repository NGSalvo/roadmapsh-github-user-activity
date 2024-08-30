package main

import (
	"flag"
	"fmt"
	"github-user-activity/services"
	"log"
)

func main() {
	githubService := services.NewGithub()

	userActivityCommand := flag.String("username", "kamranahmedse", "GitHub username")
	flag.Parse()

	events, err := githubService.GetRecentActivity(*userActivityCommand)
	if err != nil {
		log.Fatal(err)
	}
	for index, event := range events {
		fmt.Printf("---------- Event %d ----------\n", index+1)
		fmt.Println(event)
	}
}
