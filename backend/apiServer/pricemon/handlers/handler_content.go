package handlers

import (
	"pricemon/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("content", contentAll)
	groupApi.GET("content/:id",  contentOne)
	groupApi.POST("content",  contentCreate)
	groupApi.PATCH("content",  contentUpdate)
	groupApi.DELETE("content/:id",  contentDelete)
}
//All
func contentAll(c *gin.Context) {
	mdl := models.Content{}
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
//One
func contentOne(c *gin.Context) {
	var mdl models.Content
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
//Create
func contentCreate(c *gin.Context) {
	var mdl models.Content
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
//Update
func contentUpdate(c *gin.Context) {
	var mdl models.Content
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
//Delete
func contentDelete(c *gin.Context) {
	var mdl models.Content
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
