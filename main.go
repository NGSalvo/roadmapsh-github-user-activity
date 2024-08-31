package main

import (
	"github-user-activity/services"
	"os"
)

func main() {
	services.NewCommandLine().Run()
	os.Exit(0)
}
