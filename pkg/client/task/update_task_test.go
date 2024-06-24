package task

import (
	"fmt"
	"testing"
)

func TestCheckLastestVersion(t *testing.T) {
	version := getLatestVersion()
	fmt.Println(version)
}
