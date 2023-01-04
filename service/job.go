package service

import (
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/sjxiang/op-panel/pkg/helper"
)


var Cron = cron.New()

// 定时任务列表
func GetJobs(ctx *gin.Context) {
	var res []map[string]interface{}

	for _, e := range Cron.Entries() {
		res = append(res, map[string]interface{}{
			"id": e.ID,
			"next": e.Next,
		})
	}

	ctx.JSON(http.StatusOK, res)
}



// 创建定时任务
func AddJob(ctx *gin.Context) {
	var payload struct {
		Cron string `json:"cron"`
		Exec string `json:"exec"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	
	eid, err := Cron.AddFunc(payload.Cron, func() {
		ExecuteTask(payload.Exec)
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": eid,
	})
}



// 删除定时任务
func DeleteJob(ctx *gin.Context) {
	id := ctx.Param("id")
	eid, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	Cron.Remove(cron.EntryID(eid))
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}


// e.g. ls -al
func ExecuteTask(execCmd string) {
	execParts := strings.Split(execCmd, " ")
	execName := execParts[0]
	execParams := strings.Join(execParts[1:], " ")

	cmd := exec.Command(execName, execParams)
	if err := cmd.Run(); err != nil {
		helper.Fatal(err)
	}
}