package utils

import (
	"fmt"
	"os/exec"
)

func HostAlive(host string) bool {
	// 要检查的主机地址

	// 执行Ping命令
	cmd := exec.Command("ping", "-c", "1", host)
	err := cmd.Run()

	// 判断主机是否存活
	if err != nil {
		fmt.Printf("%s is not alive\n", host)
		return true
	} else {
		fmt.Printf("%s is alive\n", host)
		return false
	}
}
