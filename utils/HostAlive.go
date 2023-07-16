package utils

import (
	"fmt"
	"net"
	"sync"
)

func HostAlive() {
	// 定义需要检测的主机IP地址列表
	hosts := []string{"113.30.191.68"}

	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}

	// 使用协程并发检测主机是否存活
	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()

			_, err := net.Dial("tcp", host+":80")
			if err != nil {
				fmt.Printf("%s 不可达\n", host)
				return
			}

			fmt.Printf("%s 存活\n", host)
		}(host)
	}

	// 等待所有协程执行完成
	wg.Wait()
}
