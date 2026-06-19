package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"rightwatch/config"
	"rightwatch/models"
	"rightwatch/services"
)

func init() {
	groupApi.POST("check-list/:id/capture", captureCreate)
	groupApi.GET("check-list/:id/captures", captureList)
}

type captureReq struct {
	CaptureType string `json:"capture_type" form:"capture_type"` // list|detail|takedown; defaults to "detail"
	TriggerBy   string `json:"trigger_by"   form:"trigger_by"`   // auto|manual; defaults to "manual"
}

// captureCreate triggers a headless screenshot for the given check_list item,
// persists the JPEG under the configured captures path, and records the row in DB.
func captureCreate(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}

	var req captureReq
	_ = c.ShouldBindJSON(&req)
	if req.CaptureType == "" {
		req.CaptureType = "detail"
	}
	if req.TriggerBy == "" {
		req.TriggerBy = "manual"
	}

	postIdx, _, _, _, err := models.LoadCaptureTarget(id)
	if handleError(c, err) {
		return
	}

	pageURL := fmt.Sprintf("https://ondisk.co.kr/pop.php?sm=bbs_info&idx=%s", postIdx)

	cfg := config.GetOndiskConfig()
	jpegBytes, err := services.CaptureURL(pageURL, cfg)
	if handleError(c, err) {
		return
	}

	// Fetch current check_list status for the status_at field.
	cl := &models.CheckList{Id: id}
	clRow, err := cl.One()
	if handleError(c, err) {
		return
	}

	now := time.Now()
	cap := &models.Capture{
		CheckListId: id,
		CaptureType: req.CaptureType,
		FilePath:    "", // filled in after Create() assigns Id
		PageUrl:     pageURL,
		StatusAt:    clRow.Status,
		TriggerBy:   req.TriggerBy,
		CapturedAt:  &now,
	}
	if err := cap.Create(); handleError(c, err) {
		return
	}

	savePath := capturesPath()
	if err := os.MkdirAll(savePath, 0755); handleError(c, err) {
		return
	}
	filePath := filepath.Join(savePath, fmt.Sprintf("%d.jpg", cap.Id))
	if err := os.WriteFile(filePath, jpegBytes, 0644); handleError(c, err) {
		return
	}
	if err := cap.SaveFilePath(filePath); handleError(c, err) {
		return
	}
	cap.FilePath = filePath

	jsonData(c, cap)
}

// captureList returns all captures for a check_list item, oldest first.
func captureList(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}

	list, err := models.CapturesByCheckList(id)
	if handleError(c, err) {
		return
	}
	jsonData(c, list)
}

func capturesPath() string {
	p := viper.GetString("capture.path")
	if p == "" {
		p = "./static/captures/"
	}
	return p
}
