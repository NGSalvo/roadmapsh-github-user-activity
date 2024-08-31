package services

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	getRecentActivitySubCommand := flag.NewFlagSet("github-activity", flag.ExitOnError)
	userActivity := getRecentActivitySubCommand.String("username", "kamranahmedse", "GitHub username")
	getRecentActivitySubCommand.Parse(os.Args[2:])

	events, err := c.githubService.GetRecentActivity(*userActivity)
	if err != nil {
		return err
	}

	for index, event := range events {
		fmt.Printf("---------- Event %d ----------\n", index+1)
		fmt.Println(event)
	}

	return nil
}
