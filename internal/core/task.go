package core

import (
	"gitlab-bot/internal"
	"gitlab-bot/internal/gitlab"
	"log"
	"sync"
	"time"
)

type AutoMergeTask struct {
	taskConfig *internal.AutoMergeTaskConfiguration
	amr        internal.AutoMergeRequest
	wg         *sync.WaitGroup
}

func NewAutoMergeTask(projects *internal.AutoMergeProject, global *internal.GlobalConfiguration, wg *sync.WaitGroup) *AutoMergeTask {
	taskConfig := internal.NewAutoMergeTaskConfiguration(projects, global)
	amr := gitlab.NewAutoMergeRequest()
	return &AutoMergeTask{
		taskConfig: taskConfig,
		amr:        amr,
		wg:         wg,
	}
}

func (t *AutoMergeTask) Run() {
	err := t.amr.Init(t.taskConfig)
	if err != nil {
		log.Printf("project verify faild, project:%s, error:%s", t.taskConfig.MergeProjects.Name, err)
		t.wg.Done()
	}
	for {
		time.Sleep(time.Second * time.Duration(t.taskConfig.MergeProjects.CheckInterval))
		t.amr.MergeRequest()
	}
}

type AutoCreateMrTask struct {
	taskConfig *internal.AutoCreateMergeRequestTaskConfiguration
	acmr       internal.AutoCreateMergeRequest
	wg         *sync.WaitGroup
}

func NewAutoCreateMrTask(project *internal.AutoCreateMergeProject, global *internal.GlobalConfiguration, wg *sync.WaitGroup) *AutoCreateMrTask {
	taskConfig := internal.NewAutoCreateMergeRequestTaskConfiguration(project, global)
	acmr := gitlab.NewAutoCreateMergeRequest()
	task := &AutoCreateMrTask{
		taskConfig: taskConfig,
		acmr:       acmr,
		wg:         wg,
	}
	return task
}

func (t *AutoCreateMrTask) Run() {
	err := t.acmr.Init(t.taskConfig)
	if err != nil {
		log.Printf("project verify faild, project:%s, error:%s", t.taskConfig.Project.Name, err)
		t.wg.Done()
	}
	t.acmr.CreateMergeRequest()
}
