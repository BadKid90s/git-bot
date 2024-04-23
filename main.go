package main

import (
	"fmt"
	"github.com/kardianos/service"
	"gitlab-bot/cmd"
	"gitlab-bot/logger"
	"os"
)

func main() {

	srvConfig := &service.Config{
		Name:        "Gitlab-bot",
		DisplayName: "Gitlab-bot",
	}
	prg := &cmd.Program{}
	s, err := service.New(prg, srvConfig)
	if err != nil {
		logger.Log.Errorln(err.Error())
	}

	errs := make(chan error, 5)
	logg, err := s.Logger(errs)
	cmd.LOGGER = logg

	if err != nil {
		logger.Log.Errorln(err.Error())
		os.Exit(1)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				logger.Log.Errorln(err.Error())
			}
		}
	}()

	if len(os.Args) > 1 {
		serviceAction := os.Args[1]
		switch serviceAction {
		case "install":
			err := s.Install()
			if err != nil {
				logger.Log.Errorln(fmt.Sprintf("安装服务失败: %s", err.Error()))
			} else {
				logger.Log.Infoln("安装服务成功")
			}
			return
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				logger.Log.Errorln(fmt.Sprintf("卸载服务失败: %s", err.Error()))
			} else {
				logger.Log.Infoln("卸载服务成功")
			}
			return
		case "start":
			err := s.Start()
			if err != nil {
				logger.Log.Errorln(fmt.Sprintf("运行服务失败: %s", err.Error()))
			} else {
				logger.Log.Infoln("运行服务成功")
			}
			return
		case "stop":
			err := s.Stop()
			if err != nil {
				logger.Log.Errorln(fmt.Sprintf("停止服务失败: %s", err.Error()))
			} else {
				logger.Log.Infoln("停止服务成功")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Log.Errorln(err.Error())
	}
}
