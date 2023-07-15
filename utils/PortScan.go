package utils

import (
	"fmt"
	"net"
	"time"
)

func PortScan() {
	ip := "127.0.0.1" // 指定IP地址
	startPort := 1    // 起始端口号
	endPort := 1024   // 结束端口号
	fmt.Printf("Scanning ports %d to %d on %s\n", startPort, endPort, ip)

	for port := startPort; port <= endPort; port++ {
		target := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", target, time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
}
