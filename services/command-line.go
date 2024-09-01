package services

import (
	"flag"
	"fmt"
	"github-user-activity/models"
	"log"
	"os"
	"strings"
)

type (
	CommandLine interface {
		Run()
	}

	commandLine struct {
		githubService GithubService
	}
)

func NewCommandLine() CommandLine {
	return &commandLine{
		githubService: NewGithub(),
	}
}

func (c *commandLine) Run() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: github-activity <username>")
	}

	switch os.Args[1] {
	case "github-activity":
		err := c.getRecentActivityCommand()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid command. Expected 'github-activity'")
	}
}

func (c *commandLine) getRecentActivityCommand() error {
	var events []*models.Event
	var err error
	getRecentActivitySubCommand := flag.NewFlagSet("github-activity", flag.ExitOnError)
	userActivity := getRecentActivitySubCommand.String("username", "kamranahmedse", "GitHub username")
	eventType := getRecentActivitySubCommand.String("type", "PushEvent", "Event type")
	getRecentActivitySubCommand.Parse(os.Args[2:])

	typeFlag := false
	for _, arg := range os.Args[2:] {
		if strings.Contains(arg, "type") {
			typeFlag = true
		}
	}

	if !typeFlag {
		events, err = c.githubService.GetRecentActivity(*userActivity)
		if err != nil {
			return err
		}
	} else {
		events, err = c.githubService.GetRecentActivityByType(*userActivity, models.EventType(*eventType))
		if err != nil {
			return err
		}
	}

	if len(events) == 0 {
		fmt.Println("No events found")
		return nil
	}

	for index, event := range events {
		fmt.Printf("---------- Event %d ----------\n", index+1)
		fmt.Println(event)
	}

	return nil
}
