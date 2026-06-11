package tasks

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"rightwatch/matcher"
	"rightwatch/models"
)

// MatchingTask 스케줄러에서 호출되는 증감분 매칭 태스크.
// 마지막 완료된 매칭 crawler_job의 end_time 이후 post만 대상으로 매칭을 실행한다.
func MatchingTask() {
	db := models.GetDB()

	job := startMatchingJob(db)

	since := lastMatchingTime(db, job.Id)
	results, err := matcher.RunMatchingSince(db, since)
	if err != nil {
		logrus.WithError(err).Error("MatchingTask: RunMatchingSince failed")
		endMatchingJob(db, job, 0, 0, "error:"+err.Error())
		return
	}

	inserted := insertMatches(db, results)
	logrus.Infof("MatchingTask: matched=%d inserted=%d since=%s", len(results), inserted, since.Format(time.RFC3339))
	endMatchingJob(db, job, len(results), inserted, fmt.Sprintf("matching:matched=%d inserted=%d", len(results), inserted))
}

// startMatchingJob crawler_job 레코드를 생성하고 반환한다.
func startMatchingJob(db *gorm.DB) models.CrawlerJob {
	now := time.Now()
	job := models.CrawlerJob{
		StartTime: &now,
		Result:    "matching:running",
	}
	_ = job.Create()
	return job
}

// endMatchingJob crawler_job을 완료 상태로 업데이트한다.
func endMatchingJob(db *gorm.DB, job models.CrawlerJob, matched, inserted int, result string) {
	endTime := time.Now()
	job.EndTime = &endTime
	job.Result = result
	_ = job.Update()
}

// lastMatchingTime 마지막으로 완료된 매칭 job의 end_time을 반환한다.
// 이전 job이 없으면 zero time을 반환해 전체 스캔을 유도한다.
// currentJobID는 방금 생성한 job을 제외하기 위해 사용한다.
func lastMatchingTime(db *gorm.DB, currentJobID uint) time.Time {
	var job models.CrawlerJob
	err := db.Where("result LIKE 'matching:%' AND result != 'matching:running' AND end_time IS NOT NULL AND id != ?", currentJobID).
		Order("end_time DESC").First(&job).Error
	if err != nil || job.EndTime == nil {
		return time.Time{}
	}
	return *job.EndTime
}

// insertMatches MatchResult 목록을 check_list에 적재하고 성공 건수를 반환한다.
func insertMatches(db *gorm.DB, results []matcher.MatchResult) int {
	inserted := 0
	for _, r := range results {
		cl := models.CheckList{
			ContentId: r.ContentID,
			PostId:    r.PostID,
			PostIdx:   r.PostIdx,
			PostTxt:   r.PostTxt,
			Status:    0,
		}
		if err := cl.Create(); err == nil {
			inserted++
		}
	}
	return inserted
}
