package core

import (
	"github.com/robfig/cron"
	"gitlab-bot/internal"
	"gitlab-bot/logger"
	"golang.org/x/net/context"
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
func (b *GitLabBot) Start() {
	logger.Log.WithModule("bot").Infoln("bot starting .")

	config, err := initConfig(b.configFile)
	if err != nil {
		logger.Log.WithModule("bot").Errorln(err.Error())
		return
	}
	logger.Log.WithModule("bot").Infoln("bot config parse success.")
	b.cfg = config

	b.runMrTasks()
	b.runAutoCreateMrTasks()
	logger.Log.WithModule("bot").Infoln("bot start success.")

	for {
		select {
		case <-b.ctx.Done():
			// 如果context被取消，则停止定时任务
			logger.Log.Infoln("定时任务被停止")
			return
			//default:
			//	time.Sleep(time.Second * 1)
			//	log.Println("select")
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
			logger.Log.WithModule("bot").Errorln(err.Error())
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
			logger.Log.WithModule("bot").Errorln(err.Error())
			continue
		}
		err = c.AddFunc(project.CreateTime, func() {
			task.Run()
		})
		if err != nil {
			logger.Log.WithModule("bot").Errorf("task runing error, error: %s \n", err)
			b.wg.Done()
		}
	}
	// 启动 Cron 定时任务
	c.Start()
}

func (b *GitLabBot) SetConfig(configFile *string) {
	b.configFile = configFile
}
