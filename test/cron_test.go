package test

import (
	"testing"

	"github.com/robfig/cron/v3"

)


func TestCron(t *testing.T) { 
	
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("*/2 * * * * *", func() {  // 每 2 秒执行一次
		t.Log("Run")
	})
	if err != nil {
		t.Fatal(err)
	}
	c.Start()
	defer c.Stop()

	select {}
}