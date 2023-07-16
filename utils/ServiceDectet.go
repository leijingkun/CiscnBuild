package utils

import (
	"fmt"
	"net"
)

func ServiceDetect(ip string, port int) string {
	conn, err := net.Dial("tcp", fmt.Sprintf(ip, port))
	if err != nil {
		fmt.Printf("端口 %d 不可达\n", port)
		return "null"
	}
	defer conn.Close()

	// 发送指纹识别请求，并读取响应数据
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("端口 %d 读取数据失败\n", port)
		return "null"
	}

	// 根据响应数据的不同，识别出端口的指纹，并输出指纹识别结果
	if string(buf[:n]) == "HTTP/1.0 200 OK\r\n\r\n" {
		return "apache"
	} else if string(buf[:n]) == "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.3\r\n" {
		return "ssh"
	} else if string(buf[:n]) == "SMTP\r\n" {
		return "smtp"
	} else {
		return "unknown"
	}
}
