package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	// "github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// SystemState 获取系统状态
func SystemState(ctx *gin.Context) {
	var (
		cpuUsedPercent float64
		memUsedPercent float64
		// diskUsed       uint64
		// diskUsedPercent float64	
	)

	// cpu	
	cpuPercents, _ := cpu.Percent(time.Second, true)
	for _, percent := range cpuPercents {
		cpuUsedPercent += percent
	}
	cpuUsedPercent /= float64(len(cpuPercents))

	// mem
	vms, _ := mem.VirtualMemory()
	memUsedPercent = vms.UsedPercent

	// disk （考虑过权限嘛 permission denied）
	// partitions, _ := disk.Partitions(true)
	// for _, partition := range partitions {
	// 	us, err := disk.Usage(partition.Mountpoint)
	// 	fmt.Printf("%+v", err)
	// 	diskUsed += us.Used
	// }
	// allUsage, _ := disk.Usage("/")  
	// diskUsedPercent = float64(diskUsed) / float64(allUsage.Total) * 100  // 所有可用的磁盘


	ctx.JSON(http.StatusOK, gin.H{
		"cpu": fmt.Sprintf("%.2f", cpuUsedPercent),
		"mem": fmt.Sprintf("%.2f", memUsedPercent),
		// "disk": fmt.Sprintf("%.2f", diskUsedPercent),
	})
}