package service

import (
	"time"
)

type CronModel struct {
	Id         uint  `gorm:"auto_increment;primary_key"`
	NotifyMode uint8 `gorm:"not null;default:1"`
	Delete     uint8 `gorm:"not null;default:0"`
	Overleap   uint8 `gorm:"not null;default:0"`

	Hostname string `gorm:"size:500;not null;default:''"`
	Expr     string `gorm:"size:200;not null;default:''"`
	Shell    string `gorm:"size:4000;not null;default:''"`
	Comment  string `gorm:"size:500;not null;default:''"`
	Contact  string `gorm:"size:100;not null;default:''"`
	Notify   string `gorm:"size:100;not null;default:''"`

	CreateAt time.Time `gorm:"default:current_timestamp;not null"`
	UpdateAt time.Time `gorm:"default:'0000-00-00 00:00:00';not null"`
}

type CronExecLogModel struct {
	Id       uint      `gorm:"auto_increment;primary_key"`
	CronId   uint      `gorm:"not null;default:0"`
	Code     int       `gorm:"not null;default:0"`
	Result   string    `gorm:"size:100000;not null;default:''"` // longtext
	CreateAt time.Time `gorm:"default:current_timestamp;not null"`
}

func (CronModel) TableName() string {
	return "dcron_crons"
}

func (CronExecLogModel) TableName() string {
	return "dcron_logs"
}
