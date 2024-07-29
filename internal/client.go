package internal

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v55/github"
)

func Client(token, url string) (*github.Client, error) {
	if url == "" {
		return github.NewClient(nil).WithAuthToken(token), nil
	}
	sanitizedURL := fmt.Sprintf("https://%s", strings.TrimSuffix(url, "/"))
	return github.NewClient(nil).WithAuthToken(token).WithEnterpriseURLs(sanitizedURL, sanitizedURL)
}
