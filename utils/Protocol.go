package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"strings"
	"time"
)

// "ssh", "http", "https", "rtsp", "ftp", "telnet", "amap", "mongodb", "redis", "mysql"

// 根据端口和IP确定协议
func ProtocolDetect(ip string, port int) string {
	//conn, err := net.Dial("tcp", fmt.Sprintf(113.30.191.68, 15672))
	conn, err := net.DialTimeout("tcp", ip+":"+fmt.Sprint(port), 5*time.Second)
	//fmt.Println(ip + ":" + fmt.Sprint(port))
	if err != nil {
		fmt.Printf("端口 %d 不可达,协议探测失败\n", port)
		fmt.Println(err)
		return "null"
	}
	defer conn.Close()
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("端口 %d 读取数据失败\n", port)
		return "null"
	}

	for _, protocol := range config.Protocol_list {
		//fmt.Println(protocol)
		//fmt.Println(string(buf[:n]))
		if strings.Contains(strings.ToLower(string(buf[:n])), strings.ToLower(protocol)) {
			fmt.Printf("%s: Protocol detected\n", protocol)
			return protocol
		}
	}

	return "unkown"
}
