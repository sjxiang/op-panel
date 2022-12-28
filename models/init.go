package models

import (
	
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/logger"

	"github.com/sjxiang/op-panel/pkg/constants"	
)

var DB *gorm.DB

func NewDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,  // 解决单数表名，user
		},
		Logger: logger.Default.LogMode(logger.Info),  // SQL 语句打印输出
	})

	if err != nil {
		panic("[OPEN DB ERROR]:#" + err.Error())
	}

	migrate()
}


func migrate() {

	err := DB.AutoMigrate(&ConfigBasic{}, &TaskBasic{})
	if err != nil {
		panic("[MIGRATE ERROR]:#" + err.Error())
	}
	
}

