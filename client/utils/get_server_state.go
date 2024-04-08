package utils

import (
	"common/model"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"time"
)

func GetServerState() model.ServerState {
	var state model.ServerState

	// 获取 CPU 使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("获取 CPU 使用率失败: %v\n", err)
	} else {
		if len(cpuPercent) > 0 {
			state.CPU = cpuPercent[0]
		}
	}

	// 获取内存使用情况
	memStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("获取内存使用情况失败: %v\n", err)
	} else {
		state.MemUsed = memStat.Used
	}

	// 获取交换分区使用情况
	swapStat, err := mem.SwapMemory()
	if err != nil {
		fmt.Printf("获取交换分区使用情况失败: %v\n", err)
	} else {
		state.SwapUsed = swapStat.Used
	}

	// 获取磁盘使用情况
	partitions, e := disk.Partitions(false)
	if e != nil {
		fmt.Printf("获取磁盘信息失败: %v\n", e)
	} else {
		for _, partition := range partitions {
			usage, _ := disk.Usage(partition.Mountpoint)
			state.DiskUsed += usage.Used
		}
	}

	// 获取总网络流量
	netStat, err := net.IOCounters(false)
	if err != nil {
		fmt.Printf("获取网络流量失败: %v\n", err)
	} else if len(netStat) > 0 {
		allNetStat := netStat[0]
		state.NetInTransfer = allNetStat.BytesRecv
		state.NetOutTransfer = allNetStat.BytesSent

		// 获取网络速度
		up, down := getNetSpeed(allNetStat)

		state.NetInSpeed = up
		state.NetOutSpeed = down
	}

	// 获取系统负载
	loadStat, err := load.Avg()
	if err != nil {
		fmt.Printf("获取系统负载失败: %v\n", err)
	} else {
		state.Load1 = loadStat.Load1
		state.Load5 = loadStat.Load5
		state.Load15 = loadStat.Load15
	}

	// 获取TCP和UDP连接数
	tcpConnections, _ := net.Connections("tcp")
	udpConnections, _ := net.Connections("udp")
	state.TcpConnCount = len(tcpConnections)
	state.UdpConnCount = len(udpConnections)

	// 获取进程数
	processes, err := process.Pids()
	if err != nil {
		fmt.Printf("获取进程数失败: %v\n", err)
	} else {
		state.ProcessCount = len(processes)
	}

	state.Uptime, _ = host.Uptime()

	return state
}

// 计算当前请求的速度
func getNetSpeed(prevNetState net.IOCountersStat) (uint64, uint64) {
	var up, down uint64

	// 延时1s获取
	time.Sleep(time.Second)

	netStat, err := net.IOCounters(false)
	if err != nil {
		fmt.Printf("获取网络流量失败: %v\n", err)
		return 0, 0
	}

	nowNetStat := netStat[0]

	up = nowNetStat.BytesRecv - prevNetState.BytesRecv
	down = nowNetStat.BytesSent - prevNetState.BytesSent

	return up, down
}
