package handlers

import (
	"pricemon/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("product", productAll)
	groupApi.GET("product/:id",  productOne)
	groupApi.POST("product",  productCreate)
	groupApi.PATCH("product",  productUpdate)
	groupApi.DELETE("product/:id",  productDelete)
}
//All
func productAll(c *gin.Context) {
	mdl := models.Product{}
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
func productOne(c *gin.Context) {
	var mdl models.Product
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
func productCreate(c *gin.Context) {
	var mdl models.Product
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
func productUpdate(c *gin.Context) {
	var mdl models.Product
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
func productDelete(c *gin.Context) {
	var mdl models.Product
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
