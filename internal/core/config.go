package core

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab-bot/internal"
	"log"
	"os"
)

var cfgDir = "./config/"
var defaultCfgName = "application.yml"

func initConfig(cfgFile *string) *internal.BotConfiguration {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != nil {
		// Use config file from the flag.
		viper.SetConfigFile(*cfgFile)
	} else {
		exists, err := folderExists(cfgDir)
		if err != nil {
			log.Printf(fmt.Sprintf("folder exists error, %s", err))
			os.Exit(1)
		}
		if !exists {
			fmt.Println("Folder does not exist")
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName(defaultCfgName)
		viper.SetConfigType("yaml") //设置文件的类型//尝试进行配置读取
	}
	var botConfig *internal.BotConfiguration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	err := viper.Unmarshal(&botConfig)
	if err != nil {
		log.Printf(fmt.Sprintf("Unmarshal error, %s", err))
		os.Exit(1)
	}
	log.Printf("parse config file success.")
	return botConfig
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
