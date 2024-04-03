package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"gitlab-bot/internal"
)

func NewAutoCreateMergeRequest() internal.AutoCreateMergeRequest {
	return &autoCreateMergeRequest{}
}

type autoCreateMergeRequest struct {
	client *gitlab.Client
}

func (r *autoCreateMergeRequest) Init(config *internal.AutoCreateMergeRequestTaskConfiguration) error {
	return nil
}
func (r *autoCreateMergeRequest) CreateMergeRequest() {
}
