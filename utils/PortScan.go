package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"sync"
)

func PortScan(ip string) map[string]interface{} {
	var (
		ipInfo = make(map[string]interface{})
		mu     sync.Mutex
	)
	//执行只一次,先把设备和蜜罐检测一下
	ipInfo["ip"] = ip
	ipInfo["deviceinfo"] = DeviceDetect(ip)
	// ipInfo["honeypot"] = HoneyPot(ip)
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
				// fmt.Println("未打开:", port)
				return
			}

			conn.Close()
			fmt.Printf("端口 %d 打开\n", port)
			//对于打开的端口执行下面这个函数,探测协议和服务
			fmt.Println("当前ip:", ip)
			mu.Lock()
			defer mu.Unlock()
			Result(ip, port, ipInfo)
			//还需要对打开的端口进行蜜罐检测,检测到则添加到字符数组
		}(port)
	}

	wg.Wait()
	fmt.Println("结束,写入文件ing")
	return ipInfo
}
