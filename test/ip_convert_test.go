package test

import (
	"testing"

	"github.com/sjxiang/op-panel/pkg/helper"
	"github.com/stretchr/testify/assert"
)


func TestIPToValue(t *testing.T) {
	var tests = []string{
		"127.0.0.1",
		"192.168.0.1",
	}
	
	for idx, ip := range tests {
		
		t.Logf("当前测试 IP：%v\n", ip)
		
		if idx == 0 {
			addrNum, err := helper.IpToValue(ip)
			
			// 检测
			assert.NoError(t, err, "转换" + ip + "时，报错")
			assert.Equal(t, uint32(2130706433), addrNum, "应返回 2130706433")
		}

		if idx == 1 {
			addrNum, err := helper.IpToValue(ip)
			assert.NoError(t, err, "转换" + ip + "时，报错")
			assert.Equal(t, uint32(3232235521), addrNum, "应返回 3232235521")
		}

	}
}