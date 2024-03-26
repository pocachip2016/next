package main

import (
	_ "pricemon/config"
	"pricemon/handlers"
	"pricemon/tasks"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
