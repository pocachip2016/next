package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "rightwatch/config"
	"rightwatch/handlers"
	"rightwatch/models"
	"rightwatch/services"
	"rightwatch/tasks"
)

func main() {
	var homoglyphs []models.Homoglyph
	if err := models.GetDB().Find(&homoglyphs).Error; err != nil {
		logrus.WithError(err).Warn("homoglyph map load failed — homoglyph normalization disabled")
	} else {
		m := make(map[rune]rune, len(homoglyphs))
		for _, h := range homoglyphs {
			from, to := []rune(h.FromChar), []rune(h.ToChar)
			if len(from) == 1 && len(to) == 1 {
				m[from[0]] = to[0]
			}
		}
		services.SetHomoglyphMap(m)
	}
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
