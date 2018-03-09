package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/outman/dcron/app/service"
)

type FormValidationCronCreate struct {
	Hostname string `form:"hostname" binding:"required,max=100"`
	Expr     string `form:"expr" binding:"required,max=200"`
	Shell    string `form:"shell" binding:"required,max=4000"`
	Comment  string `form:"comment" binding:"required,max=500"`
	Contact  string `form:"contact" binding:"required,max=100"`
}

func IndexList(c *gin.Context) {
	cronService := service.NewCronService()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Data": cronService.FetchListCronsPagination(1),
	})
}

func IndexNew(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"Hosts": service.Config.Hosts,
	})
}

func IndexCreate(c *gin.Context) {
	var formValidation FormValidationCronCreate
	if err := c.ShouldBindWith(&formValidation, binding.FormPost); err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Code":  http.StatusBadRequest,
			"Error": err,
		})
		return
	}
	cronService := service.NewCronService()
	var cron = service.CronModel{
		Hostname: formValidation.Hostname,
		Shell:    formValidation.Shell,
		Expr:     formValidation.Expr,
		Comment:  formValidation.Comment,
		Contact:  formValidation.Contact,
	}
	id := cronService.CreateCron(cron)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/index/view/%d", id))
}

func IndexDelete(c *gin.Context) {
	pid := c.PostForm("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	cronService := service.NewCronService()
	cronModel := cronService.FetchCronById(id)
	if cronModel.Id <= 0 {
		c.AbortWithStatus(404)
		return
	}

	cronService.DeleteCronById(cronModel)
	c.Redirect(http.StatusMovedPermanently, "/index")
}

func IndexView(c *gin.Context) {

	pid := c.Param("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	cronService := service.NewCronService()
	cronModel := cronService.FetchCronById(id)
	if cronModel.Id <= 0 {
		c.AbortWithStatus(404)
		return
	}

	cronExecLogService := service.NewCronExecLogSevice()
	c.HTML(http.StatusOK, "view.html", gin.H{
		"Data": cronModel,
		"Logs": cronExecLogService.FetchListCronExecLogPagination(id, 1),
	})
}
