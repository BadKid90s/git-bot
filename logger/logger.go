package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type logger struct {
	file     *os.File
	logLevel LogLevel
	mu       sync.Mutex
	module   string // 新添加的字段，用于存储模块信息
}

func newLogger(filePath string, logLevel LogLevel) (*logger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &logger{file: file, logLevel: logLevel}, nil
}
func (l *logger) Debugf(message string, a any) {
	l.Printf(DEBUG, message, a)
}

func (l *logger) Infof(message string, a any) {
	l.Printf(INFO, message, a)
}

func (l *logger) Warnf(message string, a any) {
	l.Printf(WARNING, message, a)
}

func (l *logger) Errorf(message string, a any) {
	l.Printf(ERROR, message, a)
}

func (l *logger) Debugln(message string) {
	l.Println(DEBUG, message)
}

func (l *logger) Infoln(message string) {
	l.Println(INFO, message)
}

func (l *logger) Warnln(message string) {
	l.Println(WARNING, message)
}

func (l *logger) Errorln(message string) {
	l.Println(ERROR, message)
}

func (l *logger) Println(level LogLevel, message string) {
	if level < l.logLevel {
		return
	}

	levelStr := l.buildLevelString(level)
	logMessage := fmt.Sprintf("[%s] [%s] [%4s]: %s \n", time.Now().Format("2006-01-02 15:04:05"), levelStr, l.module, message)

	fmt.Printf(logMessage)
	l.write(logMessage)
}

func (l *logger) Printf(level LogLevel, message string, a any) {
	if level < l.logLevel {
		return
	}

	message = fmt.Sprintf(message, a)
	levelStr := l.buildLevelString(level)
	logMessage := fmt.Sprintf("[%s] [%s] [%4s]: %s", time.Now().Format("2006-01-02 15:04:05"), levelStr, l.module, message)

	fmt.Printf(logMessage)
	l.write(logMessage)
}

func (l *logger) buildLevelString(level LogLevel) string {
	levelStr := ""
	switch level {
	case INFO:
		levelStr = "INFO"
	case WARNING:
		levelStr = "WARNING"
	case ERROR:
		levelStr = "ERROR"
	default:
		levelStr = "INFO"
	}
	return levelStr
}

func (l *logger) write(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.file.WriteString(message)
	if err != nil {
		fmt.Printf("Failed to write log: %v\n", err)
	}
}

func (l *logger) Close() {
	_ = l.file.Close()
}

type ModuleLogger struct {
	moduleMux sync.Mutex // 新添加的互斥锁，用于保护 module 字段
	*logger
}

func (l *ModuleLogger) WithModule(module string) *ModuleLogger {
	l.moduleMux.Lock()
	defer l.moduleMux.Unlock()
	l.module = module
	return l
}

func NewModuleLogger(filePath string, logLevel LogLevel, module string) (*ModuleLogger, error) {
	newLogger, err := newLogger(filePath, logLevel)
	if err != nil {
		fmt.Println("init logger failed")
		return nil, err
	}
	m := &ModuleLogger{
		logger: newLogger,
	}
	m.module = module // 设置 module 字段为传入的 module 参数
	return m, nil
}

var Log *ModuleLogger

func init() {
	newLogger, err := NewModuleLogger("./gitlab_bot.log", INFO, "main")
	//newLogger, err := NewModuleLogger("D:\\gitlab-bot\\gitlab_bot.log", INFO, "main")
	if err != nil {
		fmt.Println("init logger failed")
		os.Exit(1)
	}
	Log = newLogger

}
