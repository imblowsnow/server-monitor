package utils

import (
	"common/model"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"strings"
	"time"
)

func GetServerInfo() model.ServerInfo {
	var info model.ServerInfo

	hostInfo, _ := host.Info()

	info.OS = hostInfo.OS
	info.HostName = hostInfo.Hostname
	info.Platform = hostInfo.Platform
	info.PlatformVersion = hostInfo.PlatformVersion
	info.PlatformFamily = hostInfo.PlatformFamily
	info.Arch = hostInfo.KernelArch
	info.BootTime = int(hostInfo.BootTime)
	info.Virtualization = hostInfo.VirtualizationSystem

	// 获取CPU信息
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("获取CPU信息失败: %v\n", err)
	} else {
		var cpus []string
		for _, cpu := range cpuInfo {
			cpus = append(cpus, cpu.ModelName)
		}
		info.CPU = strings.Join(cpus, ",")
	}

	// 获取内存信息
	memory, e := mem.VirtualMemory()
	if e != nil {
		fmt.Printf("获取内存信息失败: %v\n", e)
	} else {
		info.MemTotal = memory.Total
	}

	// 获取交换分区信息
	swap, e := mem.SwapMemory()
	if e != nil {
		fmt.Printf("获取交换分区信息失败: %v\n", e)
	} else {
		info.SwapTotal = swap.Total
	}

	// 获取磁盘信息
	partitions, e := disk.Partitions(false)
	if e != nil {
		fmt.Printf("获取磁盘信息失败: %v\n", e)
	} else {
		for _, partition := range partitions {
			usage, _ := disk.Usage(partition.Mountpoint)
			info.DiskTotal += usage.Total
		}
	}

	// 获取系统时区
	info.Timezone, _ = time.Now().Zone()

	return info
}
