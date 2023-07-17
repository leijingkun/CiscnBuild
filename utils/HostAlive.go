package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"sync"
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

			_, err := net.Dial("tcp", host+":80")
			if err != nil {
				fmt.Printf("%s 不可达\n", host)
				ipInfo = PortScan(host)
				jw := loadOutputJSON("result.json")
				jw.Push(ipInfo)
				return
			}
			fmt.Printf("%s 存活\n", host)
			ipInfo = PortScan(host)
			jw := loadOutputJSON("result.json")
			jw.Push(ipInfo)
		}(host)
	}

	// 等待所有协程执行完成
	wg.Wait()
}
