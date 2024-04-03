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

func NewAutoMergeTasks(projects []*internal.AutoMergeProject, global *internal.GlobalConfiguration, wg *sync.WaitGroup) []internal.Task {
	var tasks []internal.Task
	for _, p := range projects {
		taskConfig := internal.NewAutoMergeTaskConfiguration(p, global)
		amr := gitlab.NewAutoMergeRequest()
		task := &AutoMergeTask{
			taskConfig: taskConfig,
			amr:        amr,
			wg:         wg,
		}
		tasks = append(tasks, task)
	}
	return tasks
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

func NewAutoCreateMrTask(projects []*internal.AutoCreateMergeProject, global *internal.GlobalConfiguration, wg *sync.WaitGroup) []internal.Task {
	var tasks []internal.Task
	for _, p := range projects {
		taskConfig := internal.NewAutoCreateMergeRequestTaskConfiguration(p, global)
		acmr := gitlab.NewAutoCreateMergeRequest()
		task := &AutoCreateMrTask{
			taskConfig: taskConfig,
			acmr:       acmr,
			wg:         wg,
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (t *AutoCreateMrTask) Run() {
}
