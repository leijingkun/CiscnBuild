package utils

import (
	"fmt"
	"net"
	"sync"
)

func DeviceDetect(ip string) []string {

	return []string{"app","app"}

}
func IsHikvision(ip string) bool {
	list := []int{8000, 37020}
	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}
	for _, port := range list {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			
			}
		}(port)
	}
}
