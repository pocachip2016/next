package handlers

import (
	"github.com/gin-gonic/gin"
	"rightwatch/config"
	"rightwatch/models"
	"rightwatch/services"
)

func init() {
	groupApi.POST("check-list/:id/transition", checkListTransition)
	groupApi.POST("check-list/:id/confirm-deletion", confirmDeletion)
}

type transitionReq struct {
	From int `json:"from" form:"from"`
	To   int `json:"to"   form:"to"`
}

func checkListTransition(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}

	var req transitionReq
	if err := c.ShouldBindJSON(&req); handleError(c, err) {
		return
	}

	if err := models.Transition(id, req.From, req.To); handleError(c, err) {
		return
	}

	if req.To == models.StatusNotified {
		cpName, cpEmail, contentTitle, postURL, err := models.LoadNotifyInfo(id)
		if handleError(c, err) {
			return
		}

		subject, body := services.BuildNotification(cpName, contentTitle, postURL)
		smtpCfg := config.GetSMTP()
		dryRun, result, mailErr := services.SendMail(smtpCfg, cpEmail, subject, body)

		notif := models.Notification{
			CheckListId: id,
			ToEmail:     cpEmail,
			Subject:     subject,
			Body:        body,
			DryRun:      boolToTinyint(dryRun),
			Result:      result,
		}
		if mailErr != nil {
			notif.Result = mailErr.Error()
		}
		_ = notif.Create()
	}

	jsonSuccess(c)
}

func confirmDeletion(c *gin.Context) {
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}

	cl := &models.CheckList{Id: id}
	clRow, err := cl.One()
	if handleError(c, err) {
		return
	}

	if clRow.Status != models.StatusNotified {
		jsonError(c, "status must be notified(1) to confirm deletion")
		return
	}

	if models.PostExists(clRow.PostIdx) {
		jsonData(c, gin.H{"deleted": false, "reason": "post still exists"})
		return
	}

	if err := models.Transition(id, models.StatusNotified, models.StatusDeleteConfirmed); handleError(c, err) {
		return
	}
	jsonData(c, gin.H{"deleted": true})
}

func boolToTinyint(b bool) int {
	if b {
		return 1
	}
	return 0
}
