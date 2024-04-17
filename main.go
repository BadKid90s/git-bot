package main

import (
	"github.com/kardianos/service"
	"gitlab-bot/cmd"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("gitlab-bot.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	srvConfig := &service.Config{
		Name:        "Gitlab-bot",
		DisplayName: "Gitlab-bot",
	}
	prg := &cmd.Program{}
	s, err := service.New(prg, srvConfig)
	if err != nil {
		log.Println(err)
	}
	if len(os.Args) > 1 {
		serviceAction := os.Args[1]
		switch serviceAction {
		case "install":
			err := s.Install()
			if err != nil {
				log.Println("安装服务失败: ", err.Error())
			} else {
				log.Println("安装服务成功")
			}
			//err = s.Run()
			//if err != nil {
			//	log.Println(err)
			//}
			return
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				log.Println("卸载服务失败: ", err.Error())
			} else {
				log.Println("卸载服务成功")
			}
			return
		case "start":
			err := s.Start()
			if err != nil {
				log.Println("运行服务失败: ", err.Error())
			} else {
				log.Println("运行服务成功")
			}
			return
		case "stop":
			err := s.Stop()
			if err != nil {
				log.Println("停止服务失败: ", err.Error())
			} else {
				log.Println("停止服务成功")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		log.Println(err)
	}
}
