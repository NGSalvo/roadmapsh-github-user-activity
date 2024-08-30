package main

import (
	"fmt"
	"github-user-activity/services"
	"log"
)

func main() {
	githubService := services.NewGithub()

	events, err := githubService.GetRecentActivity("kamranahmedse")
	if err != nil {
		log.Fatal(err)
	}
	for index, event := range events {
		fmt.Printf("---------- Event %d ----------\n", index+1)
		fmt.Println(event)
	}
}
