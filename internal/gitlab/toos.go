package gitlab

import (
	"errors"
	"fmt"
	"github.com/xanzy/go-gitlab"
)

func initClient(token string, url string) (*gitlab.Client, error) {
	c, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create client: %v", err))
	}
	return c, nil
}
