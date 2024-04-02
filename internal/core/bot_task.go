package core

import (
	"gitlab-bot/internal"
	"gitlab-bot/internal/gitlab"
	"log"
	"sync"
	"time"
)

type BotTask struct {
	taskConfig *internal.TaskConfiguration
	amr        internal.AutoMergeRequest
	wg         *sync.WaitGroup
}

func NewBotTasks(projects []*internal.AutoMergeProjects, global *internal.GlobalConfiguration, wg *sync.WaitGroup) []*BotTask {
	var tasks []*BotTask
	for _, p := range projects {
		taskConfig := &internal.TaskConfiguration{
			MergeProjects: p,
			Global:        global,
		}
		amr := gitlab.New()
		task := &BotTask{
			taskConfig: taskConfig,
			amr:        amr,
			wg:         wg,
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (t *BotTask) Run() {
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
