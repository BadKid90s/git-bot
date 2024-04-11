package cmd

import (
	"github.com/kardianos/service"
	"gitlab-bot/internal/core"
	"log"
)

var serviceConfig = &service.Config{
	Name:        "gitlab-bot",
	DisplayName: "gitlab bot",
}

var bot *core.GitLabBot
var srv service.Service

type program struct {
}

func (ss *program) Start(service.Service) error {
	log.Println("coming Start.......")
	go ss.run()
	return nil
}

func (ss *program) run() {
	log.Println("coming run.......")
	err := bot.Start()
	if err != nil {
		log.Println(err)
	}
}

func (ss *program) Stop(service.Service) error {
	log.Println("coming Stop.......")
	bot.Stop()
	return nil
}

func createService() service.Service {
	log.Println("service.Interactive()---->", service.Interactive())
	return srv
}

func startAction() error {
	// 默认 运行 Run
	err := srv.Run()
	if err != nil {
		log.Printf("service Run failed, err: %v\n", err)
		return err
	}

	return nil
}

func init() {
	bot = core.NewGitLabBot()
	bot.SetConfig(&cfgFile)
	ss := &program{}
	s, err := service.New(ss, serviceConfig)
	if err != nil {
		log.Fatalln(err)
	}
	srv = s
}
