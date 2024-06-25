package task

import (
	"fmt"
	"log"
	"os/exec"
	"server-monitor/pkg/client/constants"
	"syscall"
)

func doUpdate() {
	bashCmd := "sleep 60 && curl https://raw.githubusercontent.com/imblowsnow/server-monitor/master/deploy.sh | bash -s -- " + constants.ServerIP + " " + fmt.Sprintf("%d", constants.ServerPort) + " " + constants.ServerKey + " > update.log"
	//bashCmd := "sleep 60 && echo 'hello' > update.log"

	fmt.Println("[更新检测任务] 执行 ", bashCmd)
	// 创建一个新的进程来执行 bash 命令
	cmd := exec.Command("nohup", "bash", "-c", bashCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	// 运行命令
	err := cmd.Start()
	fmt.Println("fork cmd进程ID", cmd.Process.Pid)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
