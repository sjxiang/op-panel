package service

import (
	"log"

	"github.com/robfig/cron/v3"
	"github.com/sjxiang/op-panel/models"
	"github.com/sjxiang/op-panel/pkg/helper"
)


func Cron() {
	c := cron.New(cron.WithSeconds())
	list := make([]*models.TaskBasic, 0)
	err := models.DB.Find(&list).Error
	if err != nil {
		log.Println("[DB Error]:#" + err.Error())
	}

	for _, v := range list {
		_, err := c.AddFunc(v.Spec, func() { 
			// run shell
			helper.RunShell(v.ShellPath, v.LogPath)
		})
		if err != nil {
			log.Fatal("[CRON Error]:#" + err.Error())
		}
	}
	
	c.Start()
	defer c.Stop()

	select{}
}