package gitlab

import (
	"errors"
	"fmt"
	"github.com/jucardi/go-streams/streams"
	"github.com/xanzy/go-gitlab"
	"gitlab-bot/internal"
	"log"
)

func New() internal.AutoMergeRequest {
	return &autoMergeRequest{}
}

type autoMergeRequest struct {
	client     *gitlab.Client
	projectId  int
	taskConfig *internal.TaskConfiguration
}

func (a *autoMergeRequest) Init(config *internal.TaskConfiguration) error {
	if a.client == nil {
		client, err := a.initClient(config.Global.Token, config.Global.Url)
		if err != nil {
			return err
		}
		a.client = client
	}
	a.taskConfig = config
	name := config.MergeProjects.Name
	projectId, err := GetProjectId(a.client, name)
	if err != nil {
		return err
	}
	a.projectId = projectId
	return nil
}

func (a *autoMergeRequest) isReady() ([]*gitlab.MergeRequest, error) {
	log.Printf("start check merge request is ready")
	//查询MR
	mrs, err := ProjectMergeRequests(a.client, a.projectId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get project merge request failed, error: %v", err))
	}
	//判断数量
	if len(mrs) == 0 {
		return nil, errors.New("not found project merge request")

	}
	return mrs, nil
}

func (a *autoMergeRequest) MergeRequest() {
	mrs, err := a.isReady()
	if err != nil {
		log.Printf("merge request not ready, error: %s", err)
		return
	}
	for _, mr := range mrs {
		err := a.checkNotes(mr)
		if err != nil {
			log.Printf("checked merge request notes faild, error: %s", err)
			return
		}
		err = a.accept(mr)
		if err != nil {
			log.Printf("auto merge request faild, error: %s", err)
			return
		}
	}
}

func (a *autoMergeRequest) checkNotes(mr *gitlab.MergeRequest) error {
	//获取评论
	notes, err := MergeRequestNotes(a.client, a.projectId, mr.IID)
	if err != nil {
		return errors.New(fmt.Sprintf("get merge request commits failed, error: %v", err))
	}
	//查询符合的评论个数
	matchNum := streams.FromArray(notes).
		Filter(func(v interface{}) bool {
			n := v.(*gitlab.Note)
			return n.Body == a.taskConfig.MergeProjects.Comment
		}).
		Filter(func(v interface{}) bool {
			n := v.(*gitlab.Note)
			return streams.FromArray(a.taskConfig.MergeProjects.Reviewers).Contains(n.Author.Username)
		}).Count()

	//检查是否符合配置
	if matchNum < a.taskConfig.MergeProjects.MinReviewers {
		return errors.New("merge request insufficient number of comments")
	}
	return nil
}

func (a *autoMergeRequest) accept(mr *gitlab.MergeRequest) error {
	accept, err := MergeRequestAccept(a.client, a.projectId, mr.IID)
	if err != nil {
		return err
	}
	if accept.State != "merged" {
		return errors.New(fmt.Sprintf("accept merge request failed, error %s", accept.MergeError))
	}
	log.Printf("merge request success.")
	return nil
}

func (a *autoMergeRequest) initClient(token string, url string) (*gitlab.Client, error) {
	c, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create client: %v", err))
	}
	return c, nil
}
