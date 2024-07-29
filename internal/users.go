package internal

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v55/github"
	"github.com/urfave/cli/v2"
)

// Users is the entrypoint for the users command
func Users(c *cli.Context) error {
	token := strings.TrimSpace(c.String("token"))
	hostname := strings.TrimSpace(c.String("hostname"))
	org := strings.TrimSpace(c.String("org"))
	client, err := Client(token, hostname)
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}

	fmt.Println("Retrieving repositories...")
	var repos []*github.Repository
	options := github.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	index := 1
	for {
		if index%5 == 0 {
			fmt.Printf("Processing page %d...\n", index)
		}
		index++
		reposPage, resp, err := client.Repositories.ListByOrg(c.Context, org, &github.RepositoryListByOrgOptions{
			ListOptions: options,
		})
		if err != nil {
			return fmt.Errorf("error listing repos: %w", err)
		}
		repos = append(repos, reposPage...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	index = 1
	for _, repo := range repos {
		fmt.Printf("Processing repo %d/%d...\n", index, len(repos))
		index++
		collabs, _, err := client.Repositories.ListCollaborators(c.Context, org, *repo.Name, &github.ListCollaboratorsOptions{
			ListOptions: github.ListOptions{
				Page:    1,
				PerPage: 1,
			},
			Affiliation: "direct",
		})
		if err != nil {
			return fmt.Errorf("error listing collaborators: %w", err)
		}
		if len(collabs) > 0 {
			continue
		}

		teams, _, err := client.Repositories.ListTeams(c.Context, org, *repos[0].Name, &github.ListOptions{
			Page:    1,
			PerPage: 100,
		})
		if err != nil {
			return fmt.Errorf("error listing teams: %w", err)
		}
		for _, team := range teams {
			teamMembers, _, err := client.Teams.ListTeamMembersBySlug(c.Context, org, team.GetSlug(), &github.TeamListTeamMembersOptions{
				ListOptions: github.ListOptions{
					Page:    1,
					PerPage: 1,
				},
			})
			if err != nil {
				return fmt.Errorf("error listing team members: %w", err)
			}
			if len(teamMembers) > 0 {
				break
			}
		}
		fmt.Printf("Repo %s has no users\n", *repo.Name)
	}

	return nil
}
