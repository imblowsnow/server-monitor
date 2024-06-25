package task

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCheckLastestVersion(t *testing.T) {
	version := getLatestVersion()
	fmt.Println(version)
}

func TestForkCmd(t *testing.T) {
	cmd := exec.Command("ls")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
