package core

import (
	"github.com/robfig/cron"
	"gitlab-bot/internal"
	"log"
	"sync"
)

type GitLabBot struct {
	botConfig *internal.BotConfiguration
}

func NewGitLabBot() *GitLabBot {
	return &GitLabBot{
		botConfig: botConfig,
	}
}
func (b *GitLabBot) Run() {
	var wg sync.WaitGroup
	b.runMrTasks(&wg)
	b.runAutoCreateMrTasks(&wg)

	log.Printf("bot start success.")
	wg.Wait()
}

func (b *GitLabBot) runMrTasks(wg *sync.WaitGroup) {
	for _, project := range b.botConfig.AutoMergeProjects {
		task := NewAutoMergeTask(project, b.botConfig.Global, wg)
		wg.Add(1)
		go task.Run()
	}
}

func (b *GitLabBot) runAutoCreateMrTasks(wg *sync.WaitGroup) {
	c := cron.New()
	for _, project := range b.botConfig.AutoCreateMergeRequestProjects {
		task := NewAutoCreateMrTask(project, b.botConfig.Global, wg)
		wg.Add(1)
		err := c.AddFunc(project.CreateTime, func() {
			task.Run()
		})
		if err != nil {
			log.Printf("task runing error, error: %s", err)
			wg.Done()
		}
	}
	// 启动 Cron 定时任务
	c.Start()
}
