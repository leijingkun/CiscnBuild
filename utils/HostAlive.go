package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"sync"
	"time"
)

func HostAlive() {
	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}

	// 使用协程并发检测主机是否存活
	for _, host := range config.Hosts {
		wg.Add(1)
		go func(host string) {

			defer wg.Done()

			ipInfo := make(map[string]interface{})
			//如果不存活,啥也别干了
			if checkHostAlive(host) {
				fmt.Printf("%s 不可达\n", host)
				// ipInfo = PortScan(host)
				// jw := loadOutputJSON("result.json")
				// jw.Push(ipInfo)
				return
			}
			//存活
			fmt.Printf("%s 存活\n", host)
			ipInfo = PortScan(host)
			jw := loadOutputJSON("result.json")
			jw.Push(ipInfo)
		}(host)
	}

	// 等待所有协程执行完成
	wg.Wait()
}
func checkHostAlive(ip string) bool {
	timeout := time.Second

	conn, err := net.DialTimeout("ip4:icmp", ip, timeout)
	if err != nil {
		fmt.Printf("Error pinging host %s: %s\n", ip, err)
		return true
	}
	defer conn.Close()

	// fmt.Printf("Host %s is alive\n", ip)
	return false
}
