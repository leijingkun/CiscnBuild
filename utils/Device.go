package utils

import (
	"fmt"
	"net"
	"strings"
)

func DeviceDetect(ip string) []string {
	result := []string{}
	if IsHikvision(ip) != "" {
		result = append(result, IsHikvision(ip))
	}
	if IsCisco(ip) != "" {
		result = append(result, IsCisco(ip))
	}
	if IsPfsense(ip) != "" {
		result = append(result, IsPfsense(ip))
	}
	return result
}

// curl 80端口看有没有hikvision
func IsHikvision(ip string) string {
	response, err := Curl("http://"+ip+":80", map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
	if err != nil {
		return ""
	}
	//打印输出
	// fmt.Println("response", response)
	if strings.Contains(response, "Hikvision") {
		return "hivkision"
	}
	return ""
}

// 查看80端口有无pfsense
func IsPfsense(ip string) string {
	response, err := Curl("http://"+ip+":80", map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
	if err != nil {
		return ""
	}
	// fmt.Println("response", response)
	if strings.Contains(response, "Pfsense") {
		return "Pfsense"
	}
	return ""
}

// 7547,161,30005打开即为cisco
func IsCisco(ip string) string {
	for _, port := range []int{7547, 161, 30005} {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Println("未打开:", port)
			return ""
		}
		conn.Close()
		fmt.Printf("端口 %d 打开\n", port)
		//对于打开的端口直接返回cisco,我累了
		return "cisco"
	}
	return ""
}

// 5001端口打开即为synology
func IsSynology(ip string) string {
	response, err := Curl("http://"+ip+":5001", map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
	if err != nil {
		return ""
	}
	// fmt.Println("response", response)
	if strings.Contains(response, "Synology") {
		return "Pfsense"
	}
	return ""
}
