package utils

import (
	"fmt"
	"net"
)

// 返回字段"service_app": ["Apache/2.2.15", "Joomla/N"] 需要返回服务与版本
// 版本号精确度大于等于答案为真，小于答案为假
func ServiceDetect(ip string, port int) []string {
	conn, err := net.Dial("tcp", fmt.Sprintf(ip, port))
	if err != nil {
		fmt.Printf("端口 %d 不可达\n", port)
		return []string{"aaa", "bbb"}
	}
	defer conn.Close()

	// 发送指纹识别请求，并读取响应数据
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("端口 %d 读取数据失败\n", port)
		return []string{"aaa", "bbb"}
	}
	fmt.Println(n)
	return []string{"aaa", "bbb"}
}
