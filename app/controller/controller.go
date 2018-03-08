package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/outman/dcron/app/service"
)

func IndexList(c *gin.Context) {
	cronService := service.NewCronService()
	c.JSON(200, gin.H{
		"message": cronService.FetchListCronsPagination(1),
		"one":     cronService.FetchCronById(0),
		"ids":     cronService.FetchListCrons(),
	})
}

func IndexCreate(c *gin.Context) {

}

func IndexEdit(c *gin.Context) {

}

func IndexDelete(c *gin.Context) {

}

func IndexView(c *gin.Context) {

}

func IndexLog(c *gin.Context) {

}
