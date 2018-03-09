package controller

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/outman/dcron/app/service"
	cron "gopkg.in/robfig/cron.v2"
)

type dcron struct {
	dcronCurrentCronMapEntry map[string]map[cron.EntryID]uint
	dcronController          *cron.Cron
}

func DCronNew() *dcron {
	return &dcron{
		// without mutex lock
		dcronCurrentCronMapEntry: make(map[string]map[cron.EntryID]uint),
		dcronController:          cron.New(),
	}
}

// Interval check crontab
func (d *dcron) DCronBootInterval() {
	cronService := service.NewCronService()
	ids := cronService.FetchListCrons()
	if len(ids) > 0 {
		for _, id := range ids {
			cronModel := cronService.FetchCronById(id)
			if cronModel.Id <= 0 {
				continue
			}
			d.dcronProcessing(cronModel)
		}
	}
}

// Hash one record
func (d *dcron) dcronHashCronModel(cm service.CronModel) string {
	hash := md5.New()
	io.WriteString(hash, strconv.Itoa(int(cm.Id)))
	io.WriteString(hash, "|")
	io.WriteString(hash, cm.Hostname)
	io.WriteString(hash, "|")
	io.WriteString(hash, cm.Expr)
	io.WriteString(hash, "|")
	io.WriteString(hash, cm.Shell)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Processing db crons to Cron service
func (d *dcron) dcronProcessing(cm service.CronModel) {
	hashCode := d.dcronHashCronModel(cm)
	hn, oerr := os.Hostname()
	if oerr != nil {
		return
	}

	hnbytes := bytes.ToLower([]byte(hn))
	cmhnBytes := bytes.ToLower([]byte(cm.Hostname))

	if bytes.Compare(hnbytes, cmhnBytes) != 0 {
		return
	}

	if val, ok := d.dcronCurrentCronMapEntry[hashCode]; ok {
		if cm.Delete == 1 {
			for eid, _ := range val {
				d.dcronController.Remove(eid)
			}
			delete(d.dcronCurrentCronMapEntry, hashCode)
		}
	} else {
		if cm.Delete == 0 {
			etyID, err := d.dcronController.AddFunc(cm.Expr, func() {
				var code = 0
				out, shellerr := exec.Command("/bin/sh", "-c", cm.Shell).CombinedOutput()
				if shellerr != nil {
					code = 1
				}
				cronExecLogService := service.NewCronExecLogSevice()
				logModel := service.CronExecLogModel{
					CronId: cm.Id,
					Code:   code,
					Result: string(out),
				}
				cronExecLogService.CreateCronExecLogModel(&logModel)
			})

			// ignore error
			if err == nil {
				var entityMap = make(map[cron.EntryID]uint)
				entityMap[etyID] = cm.Id
				d.dcronCurrentCronMapEntry[hashCode] = entityMap
			}
		}
	}
	return
}

// Start cron
func (d *dcron) DCronStart() {
	d.dcronController.Start()
}

// Stop cron
func (d *dcron) DCronStop() {
	d.dcronController.Stop()
}

// List all IDs
func (d *dcron) DCronListEntitys() []uint {
	var ids []uint
	entities := d.dcronController.Entries()
	for _, v := range entities {
		for _, entryMap := range d.dcronCurrentCronMapEntry {
			if val, ok := entryMap[v.ID]; ok {
				ids = append(ids, val)
			}
		}
	}
	return ids
}
