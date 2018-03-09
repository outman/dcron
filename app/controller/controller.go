package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/outman/dcron/app/service"
)

func IndexList(c *gin.Context) {
	cronService := service.NewCronService()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Data": cronService.FetchListCronsPagination(1),
	})
}

func IndexCreate(c *gin.Context) {

}

func IndexDelete(c *gin.Context) {
	pid := c.PostForm("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.AbortWithStatus(404)
	}

	cronService := service.NewCronService()
	cronModel := cronService.FetchCronById(id)
	if cronModel.Id <= 0 {
		c.AbortWithStatus(404)
	}

}

func IndexView(c *gin.Context) {

	pid := c.Param("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.AbortWithStatus(404)
	}

	cronService := service.NewCronService()
	cronModel := cronService.FetchCronById(id)
	if cronModel.Id <= 0 {
		c.AbortWithStatus(404)
	}

	cronExecLogService := service.NewCronExecLogSevice()
	c.HTML(http.StatusOK, "view.html", gin.H{
		"Data": cronModel,
		"Logs": cronExecLogService.FetchListCronExecLogPagination(id, 1),
	})
}
