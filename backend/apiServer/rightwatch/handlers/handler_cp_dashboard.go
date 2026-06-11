package handlers

import (
	"github.com/gin-gonic/gin"
	"rightwatch/models"
)

func init() {
	groupApi.GET("cp/dashboard", cpDashboard)
}

type cpDashboardRow struct {
	CpId     uint   `gorm:"column:cp_id"   json:"cp_id"`
	CpName   string `gorm:"column:cp_name" json:"cp_name"`
	Detected int    `gorm:"column:detected" json:"detected"`
}

type cpWebhardRow struct {
	CpId    uint `gorm:"column:cp_id"`
	Website int  `gorm:"column:website"`
	Count   int  `gorm:"column:count"`
}

type cpWebhardEntry struct {
	Website int `json:"website"`
	Count   int `json:"count"`
}

type cpDashboardItem struct {
	CpId      uint             `json:"cp_id"`
	CpName    string           `json:"cp_name"`
	Detected  int              `json:"detected"`
	ByWebhard []cpWebhardEntry `json:"by_webhard"`
}

func cpDashboard(c *gin.Context) {
	db := models.GetDB()

	var rows []cpDashboardRow
	err := db.Raw(`
		SELECT cp.id AS cp_id, cp.name AS cp_name, COUNT(cl.id) AS detected
		FROM cp
		LEFT JOIN kta_contents kc ON kc.cp_id = cp.id
		LEFT JOIN check_list cl ON cl.content_id = kc.id
		GROUP BY cp.id, cp.name
		ORDER BY detected DESC
	`).Scan(&rows).Error
	if handleError(c, err) {
		return
	}

	var webhardRows []cpWebhardRow
	err = db.Raw(`
		SELECT kc.cp_id, p.website, COUNT(cl.id) AS count
		FROM check_list cl
		JOIN kta_contents kc ON kc.id = cl.content_id
		JOIN post p ON p.id = CAST(cl.post_id AS UNSIGNED)
		WHERE kc.cp_id IS NOT NULL
		GROUP BY kc.cp_id, p.website
	`).Scan(&webhardRows).Error
	if handleError(c, err) {
		return
	}

	byWebhard := make(map[uint][]cpWebhardEntry)
	for _, r := range webhardRows {
		byWebhard[r.CpId] = append(byWebhard[r.CpId], cpWebhardEntry{Website: r.Website, Count: r.Count})
	}

	items := make([]cpDashboardItem, 0, len(rows))
	for _, r := range rows {
		entry := cpDashboardItem{
			CpId:      r.CpId,
			CpName:    r.CpName,
			Detected:  r.Detected,
			ByWebhard: byWebhard[r.CpId],
		}
		if entry.ByWebhard == nil {
			entry.ByWebhard = []cpWebhardEntry{}
		}
		items = append(items, entry)
	}

	jsonData(c, items)
}
