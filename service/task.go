package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/op-panel/models"
	"github.com/sjxiang/op-panel/pkg/constants"
	"github.com/sjxiang/op-panel/pkg/helper"
)

// 任务列表
func TaskList(ctx *gin.Context) {
	var (
		index = ctx.Query("index")
		size  = ctx.Query("size")

		tb = make([]*models.TaskBasic, 0)
		cnt int64
	)

	indexNum,_ := strconv.Atoi(index)
	sizeNum, _ := strconv.Atoi(size)

	indexNum = helper.If(indexNum == 0, 1, indexNum).(int)
	sizeNum = helper.If(sizeNum == 0, constants.PageSize, sizeNum).(int)

	err := models.DB.Model(new(models.TaskBasic)).Count(&cnt).Offset(indexNum).Limit(sizeNum).Find(&tb).Error
	if err != nil {
		log.Println("[DB Error]:#" + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "系统异常" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg": "加载成功",
		"data": gin.H{
			"count": cnt,
			"list": tb,
		},
	})

}


func TaskAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	spec := ctx.PostForm("spec")
	data := ctx.PostForm("data")

	if name == "" || spec == "" || data == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "必填参数不能为空",
		})

		return
	}

	shellName := helper.UUID()
	tb := &models.TaskBasic{
		Name: name,
		Spec: spec,
		ShellPath: constants.ShellDir + "/" + shellName + ".sh",
		LogPath: constants.LogPath + "/" + shellName + ".log",
	}


	err := models.DB.Create(tb).Error
	if err != nil {
		log.Println("[DB Error]:#" + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "系统异常" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "新增成功",
		"code": 200,
	})
}


func TaskEdit(ctx *gin.Context) {

}


func TaskDelete(ctx *gin.Context) {

}

