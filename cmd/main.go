package main

import (
	"gitlab-bot/internal/core"
)

func main() {
	bot := core.NewGitLabBot()
	bot.Run()
}
