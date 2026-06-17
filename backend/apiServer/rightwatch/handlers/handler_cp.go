package handlers

import (
	"github.com/gin-gonic/gin"
	"rightwatch/models"
)

func init() {
	groupApi.GET("cp", cpAll)
	groupApi.GET("cp/:id", cpOne)
	groupApi.POST("cp", cpCreate)
	groupApi.PATCH("cp", cpUpdate)
	groupApi.DELETE("cp/:id", cpDelete)
}

// All 저작권자 전체 조회
func cpAll(c *gin.Context) {
	mdl := models.Cp{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

// One 저작권자 단건 조회
func cpOne(c *gin.Context) {
	var mdl models.Cp
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

// Create 저작권자 생성
func cpCreate(c *gin.Context) {
	var mdl models.Cp
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

// Update 저작권자 수정
func cpUpdate(c *gin.Context) {
	var mdl models.Cp
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

// Delete 저작권자 삭제
func cpDelete(c *gin.Context) {
	var mdl models.Cp
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
