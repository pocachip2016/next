package handlers

import (
	"pricemon/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("price-attime", priceAttimeAll)
	groupApi.GET("price-attime/:id",  priceAttimeOne)
	groupApi.POST("price-attime",  priceAttimeCreate)
	groupApi.PATCH("price-attime",  priceAttimeUpdate)
	groupApi.DELETE("price-attime/:id",  priceAttimeDelete)
}
//All
func priceAttimeAll(c *gin.Context) {
	mdl := models.PriceAttime{}
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
func priceAttimeOne(c *gin.Context) {
	var mdl models.PriceAttime
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
func priceAttimeCreate(c *gin.Context) {
	var mdl models.PriceAttime
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
func priceAttimeUpdate(c *gin.Context) {
	var mdl models.PriceAttime
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
func priceAttimeDelete(c *gin.Context) {
	var mdl models.PriceAttime
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
