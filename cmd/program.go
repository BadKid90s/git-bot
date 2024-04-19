package cmd

import (
	"github.com/kardianos/service"
	"gitlab-bot/internal/core"
	"path/filepath"
)

var bot *core.GitLabBot
var LOGGER service.Logger

type Program struct{}

func (p *Program) Start(service.Service) error {
	LOGGER.Info("server start running...")
	go p.run()
	return nil
}
func (p *Program) run() {
	LOGGER.Info("server running...")
	// 具体的服务实现
	bot.Start()
}

func (p *Program) Stop(service.Service) error {
	LOGGER.Info("server stop...")
	bot.Stop()
	return nil
}

func init() {
	//configName := filepath.Join(".", "config", "gitlab-bot.yml")
	configName := filepath.Join("D:\\", "gitlab-bot.yml")

	bot = core.NewGitLabBot()
	bot.SetConfig(&configName)
}
