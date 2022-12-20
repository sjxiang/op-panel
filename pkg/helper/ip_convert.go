package helper

import (
	"encoding/binary"
	"errors"
	"net"
)


// IP 字符串和 IP 数值之间的转换

func IpToValue(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("Malformed IP address")  // 错误的 IP 地址
	}
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("Malformed IP address")
	}
	return binary.BigEndian.Uint32(ip), nil
}


func ValueToIP(val uint32) net.IP {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, val)
	ip := net.IP(bytes)
	return ip
}