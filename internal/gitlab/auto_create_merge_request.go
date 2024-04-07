package gitlab

import (
	"errors"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"gitlab-bot/internal"
	"log"
)

func NewAutoCreateMergeRequest() internal.AutoCreateMergeRequest {
	return &autoCreateMergeRequest{}
}

type autoCreateMergeRequest struct {
	client        *gitlab.Client
	projectConfig *internal.AutoCreateMergeProject
	projectId     int
	userMap       map[string]int
}

func (a *autoCreateMergeRequest) Init(config *internal.AutoCreateMergeRequestTaskConfiguration) error {
	if a.client == nil {
		token, err := config.GetToken()
		if err != nil {
			return err
		}
		url, err := config.GetUrl()
		if err != nil {
			return err
		}
		client, err := initClient(token, url)
		if err != nil {
			return err
		}
		a.client = client
	}
	a.projectConfig = config.Project
	name := config.Project.Name
	projectId, err := GetProjectId(a.client, name)
	if err != nil {
		return err
	}
	a.projectId = projectId

	users, err := GetUserInfo(a.client, projectId)
	if err != nil {
		return err
	}
	usernames := []string{
		a.projectConfig.Assignee,
	}
	usernames = append(usernames, a.projectConfig.Reviewers...)
	userMap := make(map[string]int)
	for _, username := range usernames {
		flag := false
		for _, user := range users {
			uname := user.Username
			if username == uname {
				userMap[username] = user.ID
				flag = true
				break
			}
		}
		if !flag {
			return errors.New(fmt.Sprintf("project not found user %s", username))
		}
	}
	a.userMap = userMap
	return nil
}

func (a *autoCreateMergeRequest) CreateMergeRequest() {
	log.Printf("start auto create merge request")

	compareResult, err := ProjectBranchCompare(a.client, a.projectId, a.projectConfig.TargetBranch, a.projectConfig.SourceBranch)
	if err != nil {
		log.Printf("compare branch faild, error: %s", err)
		return
	}
	if !compareResult {
		log.Printf("branch not diff, no create merge request")
		return
	}

	err = a.createMR()
	if err != nil {
		log.Printf("auto create merge request faild, error: %s", err)
		return
	}
}

func (a *autoCreateMergeRequest) createMR() error {
	assigneeId := a.userMap[a.projectConfig.Assignee]
	var reviewerIds []int
	for _, reviewer := range a.projectConfig.Reviewers {
		reviewerIds = append(reviewerIds, a.userMap[reviewer])
	}

	labels := a.projectConfig.Labels
	_, err := CreateMergeRequest(
		a.client, a.projectId, a.projectConfig.SourceBranch,
		a.projectConfig.TargetBranch, a.projectConfig.Title, a.projectConfig.Description,
		assigneeId, reviewerIds, labels,
	)
	return err
}
