package constants

import (
	"time"
)

// 常量，配置相关

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"

	MySQLDefaultDSN         = "root:123456@tcp(172.20.0.1:3306)/op-panel?charset=utf8&parseTime=True&loc=Local"
	DefaultLimit            = 10
	TokenExpireDuation      = time.Hour * 720   // 1 月
)


const PageSize int = 20
const ShellDir = "./shell"
const LogPath = "./log"
