package handlers

import (
	"pricemon/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("price-list", priceListAll)
	groupApi.GET("price-list/:id",  priceListOne)
	groupApi.POST("price-list",  priceListCreate)
	groupApi.PATCH("price-list",  priceListUpdate)
	groupApi.DELETE("price-list/:id",  priceListDelete)
}
//All
func priceListAll(c *gin.Context) {
	mdl := models.PriceList{}
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
func priceListOne(c *gin.Context) {
	var mdl models.PriceList
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
func priceListCreate(c *gin.Context) {
	var mdl models.PriceList
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
func priceListUpdate(c *gin.Context) {
	var mdl models.PriceList
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
func priceListDelete(c *gin.Context) {
	var mdl models.PriceList
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
