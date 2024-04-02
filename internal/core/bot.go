package core

import (
	"gitlab-bot/internal"
	"log"
	"sync"
)

type GitLabBot struct {
	botConfig *internal.BotConfiguration
}

func NewGitLabBot() *GitLabBot {
	botConfig := ParseConfigFile()
	log.Printf("parse config file success.")
	return &GitLabBot{
		botConfig: botConfig,
	}
}
func (b *GitLabBot) Run() {
	var wg sync.WaitGroup
	tasks := NewBotTasks(b.botConfig.AutoMergeProjects, b.botConfig.Global, nil)
	for _, task := range tasks {
		wg.Add(1)
		go task.Run()
	}
	log.Printf("bot start success.")
	wg.Wait()
}
