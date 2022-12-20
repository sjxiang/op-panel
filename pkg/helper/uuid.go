package helper

import (
	uuid "github.com/satori/go.uuid"
)


// uuid （考虑雪花算法，时间回拨）
func UUID() string {
	return uuid.NewV4().String()
}
