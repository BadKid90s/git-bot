package core

import (
	"github.com/robfig/cron"
	"gitlab-bot/internal"
	"golang.org/x/net/context"
	"log"
	"sync"
)

type GitLabBot struct {
	configFile *string
	cfg        *internal.BotConfiguration
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	c          *cron.Cron
}

func NewGitLabBot() *GitLabBot {
	// 创建一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())
	return &GitLabBot{
		ctx:    ctx,
		cancel: cancel,
		c:      cron.New(),
	}
}
func (b *GitLabBot) Start() error {
	go b.run()
	return nil
}

func (b *GitLabBot) run() {
	log.Println("bot starting .")

	b.cfg = initConfig(b.configFile)

	b.runMrTasks()
	b.runAutoCreateMrTasks()

	log.Println("bot start success.")

	for {
		select {
		case <-b.ctx.Done():
			// 如果context被取消，则停止定时任务
			log.Println("定时任务被停止")
			return
		}
	}
}

func (b *GitLabBot) Stop() {
	b.cancel()  //取消所有任务
	b.wg.Wait() //等待所有任务协程结束
	b.c.Stop()
}

func (b *GitLabBot) runMrTasks() {
	for _, project := range b.cfg.AutoMergeProjects {
		task := NewAutoMergeTask(project, b.cfg.Global, b.ctx)

		err := task.Init()
		if err != nil {
			continue
		}
		b.wg.Add(1)

		go func(t internal.Task) {
			defer b.wg.Done()
			t.Run()
		}(task)
	}
}

func (b *GitLabBot) runAutoCreateMrTasks() {
	c := cron.New()
	for _, project := range b.cfg.AutoCreateMergeRequestProjects {
		task := NewAutoCreateMrTask(project, b.cfg.Global, b.ctx)
		err := task.Init()
		if err != nil {
			continue
		}
		err = c.AddFunc(project.CreateTime, func() {
			task.Run()
		})
		if err != nil {
			log.Printf("task runing error, error: %s \n", err)
			b.wg.Done()
		}
	}
	// 启动 Cron 定时任务
	c.Start()
}

func (b *GitLabBot) SetConfig(configFile *string) {
	b.configFile = configFile
}
