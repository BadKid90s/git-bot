package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

var Log *logrus.Logger

func InitLogger() {
	// 自定义 Logger
	log := logrus.New()

	// 同时写到多个输出
	w1 := os.Stdout
	//logFilePath := filepath.Join(".", "gitlab-bot.log")
	logFilePath := filepath.Join("D:\\", "gitlab-bot.log")
	w2, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)

	defer func(w2 *os.File) {
		_ = w2.Close()
	}(w2)

	if err != nil {
		log.Errorf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}

	log.SetOutput(io.MultiWriter(w1, w2)) // io.MultiWriter 返回一个 io.Writer 对象

	// Logger 对象的属性基本都是导出字段，可以通过 SetXxx 方法修改，也可以直接赋值修改
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Formatter = &logrus.TextFormatter{
		DisableColors: false, // 控制台颜色输出
		FieldMap: logrus.FieldMap{ // 允许用户自定义默认字段的键名
			logrus.FieldKeyMsg: "message",
		},
	}

	Log = log
}
