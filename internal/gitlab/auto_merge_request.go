package gitlab

import (
	"errors"
	"fmt"
	"github.com/todocoder/go-stream/stream"
	"github.com/xanzy/go-gitlab"
	"gitlab-bot/internal"
	"gitlab-bot/logger"
)

func NewAutoMergeRequest() internal.AutoMergeRequest {
	return &autoMergeRequest{}
}

type autoMergeRequest struct {
	client     *gitlab.Client
	projectId  int
	taskConfig *internal.AutoMergeTaskConfiguration
}

func (a *autoMergeRequest) Init(config *internal.AutoMergeTaskConfiguration) error {
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
	logger.Log.Infoln("start check merge request is ready")
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
		logger.Log.Infof("merge request not ready, error: %s \n", err)
		return
	}
	for _, mr := range mrs {
		err := a.checkNotes(mr)
		if err != nil {
			logger.Log.Infof("checked merge request notes faild, error: %s \n", err)
			return
		}
		err = a.accept(mr)
		if err != nil {
			logger.Log.Infof("auto merge request faild, error: %s \n", err)
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
	matchNum := stream.Of(notes...).
		Filter(func(n *gitlab.Note) bool {
			return n.Body == a.taskConfig.MergeProjects.Comment
		}).
		Filter(func(n *gitlab.Note) bool {
			return stream.Of(a.taskConfig.MergeProjects.Reviewers...).
				AnyMatch(func(s string) bool {
					return s == n.Author.Username
				})
		}).Count()

	//检查是否符合配置
	if matchNum < int64(a.taskConfig.MergeProjects.MinReviewers) {
		return errors.New("merge request insufficient number of comments")
	}
	return nil
}

func (a *autoMergeRequest) accept(mr *gitlab.MergeRequest) error {
	accept, err := MergeRequestAccept(a.client, a.projectId, mr.IID, a.taskConfig.MergeProjects.RemoveSourceBranch)
	if err != nil {
		return err
	}
	if accept.State != "merged" {
		return errors.New(fmt.Sprintf("accept merge request failed, error %s", accept.MergeError))
	}
	logger.Log.Infof("merge request success.")
	return nil
}
