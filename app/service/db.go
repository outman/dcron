package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	pageSize = 1000
)

type pagination struct {
	Page     uint
	Total    int
	PageSize uint
	Data     interface{}
}

// db
func FetchDbConnection() *gorm.DB {
	conn, err := gorm.Open(Config.Db.Driver, Config.Db.Dsn)
	if err != nil {
		panic(err)
	}
	conn.LogMode(true)
	return conn
}

// cron records services
type cronService struct{}

func NewCronService() *cronService {
	return &cronService{}
}

func (service *cronService) FetchListCronsPagination(page uint) pagination {
	conn := FetchDbConnection()
	defer conn.Close()

	total := 0
	conn.Model(&CronModel{}).Count(&total)

	var crons []CronModel
	offset := (page - 1) * pageSize
	conn.Order("`hostname` asc, `delete` asc, `id` desc").Offset(offset).Limit(pageSize).Find(&crons)

	return pagination{Page: page, Total: total, PageSize: pageSize, Data: crons}
}

func (service *cronService) CreateCron(cron CronModel) uint {
	conn := FetchDbConnection()
	defer conn.Close()
	conn.Create(&cron)
	return cron.Id
}

func (service *cronService) DeleteCronById(cron CronModel) {
	conn := FetchDbConnection()
	defer conn.Close()
	conn.Model(&cron).UpdateColumn("delete", 1)
	return
}

func (service *cronService) FetchListCrons() []int {
	conn := FetchDbConnection()
	defer conn.Close()

	var ids []int
	var crons []CronModel
	conn.Select("id").Find(&crons).Pluck("id", &ids)
	return ids
}

func (service *cronService) FetchCronById(id int) CronModel {
	conn := FetchDbConnection()
	defer conn.Close()

	var cron CronModel
	conn.First(&cron, id)
	return cron
}

// Cron exce log service
type cronExecLogService struct {
}

func NewCronExecLogSevice() *cronExecLogService {
	return &cronExecLogService{}
}

func (service *cronExecLogService) CreateCronExecLogModel(logModel *CronExecLogModel) {
	conn := FetchDbConnection()
	defer conn.Close()
	conn.Create(&logModel)
}

func (service *cronExecLogService) FetchListCronExecLogPagination(cronId int, page uint) pagination {
	conn := FetchDbConnection()
	defer conn.Close()

	total := 0
	conn.Where("cron_id = ?", cronId).Model(&CronExecLogModel{}).Count(&total)

	var cronExecLogs []CronExecLogModel
	offset := (page - 1) * pageSize
	conn.Where("cron_id = ?", cronId).Order("id desc").Offset(offset).Limit(pageSize).Find(&cronExecLogs)

	return pagination{Page: page, Total: total, PageSize: pageSize, Data: cronExecLogs}
}
