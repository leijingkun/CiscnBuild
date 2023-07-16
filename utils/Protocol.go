package utils

import (
	"fmt"
	"net"
)

// 根据端口和IP确定协议
func ProtocolDetect(ip string, port int) string {
	conn, err := net.Dial("tcp", fmt.Sprintf(ip, port))
	if err != nil {
		fmt.Printf("端口 %d 不可达,协议探测失败\n", port)
		return "null"
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("端口 %d 读取数据失败\n", port)
		return "null"
	}

	// 判断协议类型，并输出协议识别结果
	if string(buf[:n]) == "HTTP/1.0 200 OK\r\n\r\n" {
		return "http"
	} else if string(buf[:n]) == "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.3\r\n" {
		return "ssh"
	} else if string(buf[:n]) == "SMTP\r\n" {
		return "smtp"
	} else {
		return "unkown"
	}
}
