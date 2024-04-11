package core

import (
	"fmt"
	"gitlab-bot/internal"
	"gitlab-bot/internal/gitlab"
	"golang.org/x/net/context"
	"time"
)

func NewAutoMergeTask(projects *internal.AutoMergeProject, global *internal.GlobalConfiguration, ctx context.Context) *AutoMergeTask {
	taskConfig := internal.NewAutoMergeTaskConfiguration(projects, global)
	amr := gitlab.NewAutoMergeRequest()
	return &AutoMergeTask{
		taskConfig: taskConfig,
		amr:        amr,
		ctx:        ctx,
	}
}

type AutoMergeTask struct {
	ctx        context.Context
	taskConfig *internal.AutoMergeTaskConfiguration
	amr        internal.AutoMergeRequest
}

func (t *AutoMergeTask) Init() error {
	return t.amr.Init(t.taskConfig)
}
func (t *AutoMergeTask) Run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("MR Task has stopped\n")
			return
		default:
			time.Sleep(time.Second * time.Duration(t.taskConfig.MergeProjects.CheckInterval))
			t.amr.MergeRequest()
		}
	}
}

func NewAutoCreateMrTask(project *internal.AutoCreateMergeProject, global *internal.GlobalConfiguration, ctx context.Context) *AutoCreateMrTask {
	taskConfig := internal.NewAutoCreateMergeRequestTaskConfiguration(project, global)
	acmr := gitlab.NewAutoCreateMergeRequest()
	task := &AutoCreateMrTask{
		taskConfig: taskConfig,
		acmr:       acmr,
		ctx:        ctx,
	}
	return task
}

type AutoCreateMrTask struct {
	ctx        context.Context
	taskConfig *internal.AutoCreateMergeRequestTaskConfiguration
	acmr       internal.AutoCreateMergeRequest
}

func (t *AutoCreateMrTask) Init() error {
	return t.acmr.Init(t.taskConfig)
}

func (t *AutoCreateMrTask) Run() {
	select {
	case <-t.ctx.Done():
		fmt.Printf(" CreateMR Task has stopped\n")
		return
	default:
		t.acmr.CreateMergeRequest()
	}
}
