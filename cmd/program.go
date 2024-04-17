package cmd

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/mitchellh/go-homedir"
	"gitlab-bot/internal/core"
	"log"
	"path/filepath"
)

var bot *core.GitLabBot

type Program struct{}

func (p *Program) Start(service.Service) error {
	log.Println("server running...")
	go p.run()
	return nil
}
func (p *Program) run() {
	// 具体的服务实现
	bot.Start()
}

func (p *Program) Stop(service.Service) error {
	bot.Stop()
	return nil
}

func init() {
	bot = core.NewGitLabBot()
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	configName := fmt.Sprintf("%s%sgitlab-bot.yml", dir, string(filepath.Separator))
	bot.SetConfig(&configName)
}
