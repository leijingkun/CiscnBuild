package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"sync"
)

func PortScan(ip string) {
	ipInfo["ip"] = ip
	ipInfo["deviceinfo"] = DeviceDetect(ip)
	ipInfo["honeypot"] = HoneyPot(ip)
	// 定义需要扫描的IP地址和端口号

	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}

	// 使用协程并发扫描端口
	for _, port := range config.Ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", ip, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("端口 %d 打开\n", port)
			Result(ip, port)
		}(port)
	}

	wg.Wait()
	jw := loadOutputJSON("result.json")
	jw.Push(ipInfo)
}
