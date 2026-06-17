package handlers

import (
	"github.com/gin-gonic/gin"
	"rightwatch/matcher"
	"rightwatch/models"
)

func init() {
	groupApi.POST("matching/run", matchingRun)
	groupApi.POST("matching/run-for-content/:id", matchingRunForContent)
	groupApi.GET("matching/status", matchingStatus)
}

func matchingRun(c *gin.Context) {
	results, err := matcher.RunMatching(models.GetDB())
	if handleError(c, err) {
		return
	}

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

	jsonData(c, gin.H{"matched": len(results), "inserted": inserted})
}

func matchingRunForContent(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	results, err := matcher.RunMatchingForContent(models.GetDB(), id)
	if handleError(c, err) {
		return
	}
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
	jsonData(c, gin.H{"matched": len(results), "inserted": inserted, "content_id": id})
}

func matchingStatus(c *gin.Context) {
	type statusRow struct {
		Status int `gorm:"column:status"`
		Count  int `gorm:"column:count"`
	}
	var rows []statusRow
	if err := models.GetDB().Model(&models.CheckList{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&rows).Error; err != nil {
		handleError(c, err)
		return
	}

	byStatus := make(map[int]int, len(rows))
	total := 0
	for _, row := range rows {
		byStatus[row.Status] = row.Count
		total += row.Count
	}

	jsonData(c, gin.H{"total": total, "by_status": byStatus})
}
