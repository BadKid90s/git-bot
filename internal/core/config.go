package core

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gitlab-bot/internal"
	"os"
)

var cfgDir = "./config/"
var defaultCfgName = "application.yml"

func initConfig(cfgFile *string) (*internal.BotConfiguration, error) {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != nil {
		// Use config file from the flag.
		viper.SetConfigFile(*cfgFile)
	} else {
		exists, err := folderExists(cfgDir)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("config file folder exists error, %s", err))
		}
		if !exists {
			str, _ := os.Getwd()
			return nil, errors.New(fmt.Sprintf("config file folder does not exist,path :%s", str))
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName(defaultCfgName)
		viper.SetConfigType("yaml") //设置文件的类型//尝试进行配置读取
	}
	var botConfig *internal.BotConfiguration

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("Can't read config: %s \n", err))
	}

	err := viper.Unmarshal(&botConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unmarshal error, %s", err))
	}
	return botConfig, nil
}

// FolderExists 判断指定路径的文件夹是否存在
func folderExists(folderPath string) (bool, error) {
	fileInfo, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}
