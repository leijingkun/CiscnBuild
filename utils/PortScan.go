package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"sync"
)

func PortScan(ip string) map[string]interface{} {
	var ipInfo = make(map[string]interface{})
	//执行只一次,先把设备和蜜罐检测一下
	ipInfo["ip"] = ip
	ipInfo["deviceinfo"] = DeviceDetect(ip)
	ipInfo["honeypot"] = HoneyPot(ip, 22)
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
			//对于打开的端口执行下面这个函数,探测协议和服务
			fmt.Println("当前ip:", ip)
			Result(ip, port, ipInfo)
		}(port)
	}

	wg.Wait()
	//当前主机的所有任务完成
	return ipInfo
}
