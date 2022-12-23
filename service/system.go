package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/sjxiang/op-panel/models"
	"github.com/sjxiang/op-panel/pkg/helper"
)

// 配合 shouldBind
type UserInfoRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}


func Login(ctx *gin.Context) {

	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if name == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "用户名或者密码不能为空",
		})

		return
	}

	cb := new(models.ConfigBasic)
	err := models.DB.Model(new(models.ConfigBasic)).Where("`key` = ?", name).First(cb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg": "用户信息未初始化",
			})

			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "数据查询异常" + err.Error(),
		})

		return
	}

	fmt.Println(cb.Key, cb.Value)

	if cb.Key != name || cb.Value != password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "用户名或密码错误",
		})

		return
	}

	
	// 生成 JWT
	token, err := helper.GenerateToken(cb.Key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "系统异常" + err.Error(),
		})

		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "登录成功",
		"data": gin.H{
			"token": "Bearer " + token,
		},
	})
}

