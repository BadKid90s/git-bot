package gitlab

import (
	"errors"
	"github.com/xanzy/go-gitlab"
	"log"
)

func GetProjectId(client *gitlab.Client, name string) (int, error) {
	opt := &gitlab.ListProjectsOptions{
		Owned: gitlab.Ptr(true),
	}
	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		log.Printf("get projects error: %v", err)
	}
	for _, project := range projects {
		if project.Name == name {
			return project.ID, nil
		}
	}
	return 0, errors.New("project not matched")
}

func ProjectMergeRequests(client *gitlab.Client, projectId int) ([]*gitlab.MergeRequest, error) {
	opt := &gitlab.ListProjectMergeRequestsOptions{
		State: gitlab.Ptr("opened"),
		//Scope: gitlab.Ptr("assigned-to-me"),
		Scope: gitlab.Ptr("assigned_to_me"),
	}
	mrs, _, err := client.MergeRequests.ListProjectMergeRequests(projectId, opt)
	if err != nil {
		return nil, err
	}
	return mrs, nil
}

func MergeRequestNotes(client *gitlab.Client, projectId int, mrId int) ([]*gitlab.Note, error) {
	opt := &gitlab.ListMergeRequestNotesOptions{}
	notes, _, err := client.Notes.ListMergeRequestNotes(projectId, mrId, opt)
	if err != nil {
		return nil, err
	}
	result := make([]*gitlab.Note, 0)
	for _, value := range notes {
		if value.System == false {
			result = append(result, value)
		}
	}
	return result, nil
}

func MergeRequestAccept(client *gitlab.Client, projectId int, mrId int) (*gitlab.MergeRequest, error) {
	opt := &gitlab.AcceptMergeRequestOptions{}
	req, _, err := client.MergeRequests.AcceptMergeRequest(projectId, mrId, opt)
	if err != nil {
		return nil, err
	}
	return req, nil
}
