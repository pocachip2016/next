package tasks

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RunTasks 스케줄 태스크를 등록하고 스케줄러를 시작한다.
// app.matching_interval_hours (기본 1) 간격으로 증감분 매칭을 실행한다.
func RunTasks() {
	interval := uint64(viper.GetInt("app.matching_interval_hours"))
	if interval == 0 {
		interval = 1
	}
	logrus.Infof("MatchingTask scheduled every %d hour(s)", interval)
	Every(interval).Hours().Do(MatchingTask)
	<-Start()
}
