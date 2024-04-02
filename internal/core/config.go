package core

import (
	"github.com/spf13/viper"
	"gitlab-bot/internal"
	"os"
)

func ParseConfigFile() *internal.BotConfiguration {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	vip := viper.New()
	vip.AddConfigPath(path + "/config") //设置读取的文件路径
	vip.SetConfigName("application")    //设置读取的文件名
	vip.SetConfigType("yaml")           //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	var botConfig *internal.BotConfiguration

	err = vip.Unmarshal(&botConfig)
	if err != nil {
		panic(err)
	}
	return botConfig
}
