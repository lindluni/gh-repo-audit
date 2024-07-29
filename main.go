package main

import (
	"fmt"
	"os"

	"github.com/lindluni/gh-repo-audit/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "gh-repo-audit",
		Usage:                "Audit GitHub repositories",
		Version:              "1.0.0",
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			{
				Name:   "users",
				Usage:  "Retrieve a list of repos with no users",
				Flags:  internal.UsersFlags(),
				Action: internal.Users,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
