package app

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/outman/dcron/app/controller"
)

func RunServer(host string, port string) {

	router := gin.Default()

	router.Static("/assets", "resource/assets")
	router.LoadHTMLGlob("resource/templates/*")

	router.GET("/", controller.IndexList)

	indexGroup := router.Group("/index")
	indexGroup.GET("/", controller.IndexList)
	indexGroup.GET("/view/:id", controller.IndexView)
	indexGroup.POST("/create", controller.IndexCreate)
	indexGroup.POST("/delete", controller.IndexDelete)

	router.Run(host + ":" + port)
}

func logEntryToFile(entry []uint) {

	path := "logs/dcron.cli.log"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	var entryArray []string
	for _, v := range entry {
		entryArray = append(entryArray, strconv.Itoa(int(v)))
	}
	log.Println("[", time.Now(), "] ID : ->|"+strings.Join(entryArray, ","))
}

func RunCrond() {
	dcron := controller.DCronNew()
	dcron.DCronBootInterval()
	dcron.DCronStart()

	logEntryToFile(dcron.DCronListEntitys())

	for {
		time.Sleep(1000 * time.Millisecond * 60)
		dcron.DCronBootInterval()
		logEntryToFile(dcron.DCronListEntitys())
	}
}
