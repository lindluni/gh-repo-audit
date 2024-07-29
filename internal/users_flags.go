package internal

import "github.com/urfave/cli/v2"

// UsersFlags returns the CLI flags for the users command
func UsersFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "hostname",
			Usage:    "GitHub API hostname",
			Required: false,
			Aliases:  []string{"i"},
		},
		&cli.StringFlag{
			Name:     "org",
			Usage:    "GitHub organization",
			Required: false,
			Aliases:  []string{"o"},
		},
		&cli.StringFlag{
			Name:     "token",
			Usage:    "GitHub API token",
			Required: true,
			Aliases:  []string{"t"},
		},
	}
}
