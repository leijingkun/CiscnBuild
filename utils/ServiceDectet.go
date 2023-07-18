package utils

import (
	"CiscnMap/config"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

//"Windows", "centos", "ubuntu", "openssh", "openssl", "wordpress", "LiteSpeed", "Jetty", "java", "node.js", "express"
//, "asp.net", "php", "Microsoft-HTTPAPI", "rabbitmp", "apache", "iis", "nginx", "micro_httpd", "openresty", "grafana", "Weblogic", "elasticsearch", "debian"

func ServiceDetect(ip string, port int) []string {
	conn, err := net.DialTimeout("tcp", ip+":"+fmt.Sprint(port), 5*time.Second)
	if err != nil {
		fmt.Printf("端口 %d 不可达\n", port)
		return nil
	}
	defer conn.Close()

	// 发送指纹识别请求，并读取响应数据
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("端口 %d 读取数据失败\n", port)
		return nil
	}

	// 根据响应数据的不同，识别出端口的指纹，并输出指纹识别结果
	fmt.Println(string(buf[:n]))
	bufLower := strings.ToLower(string(buf[:n]))

	serviceVersions := []string{}
	for _, service := range config.Server_list {
		serviceLower := strings.ToLower(service)
		if strings.Contains(bufLower, serviceLower) {
			fmt.Printf("%s: Service detected\n", service)

			escapedService := regexp.QuoteMeta(serviceLower)

			versionRegex := regexp.MustCompile(`(?i)` + escapedService + `[-_]+([\d.]+)`)
			matches := versionRegex.FindStringSubmatch(bufLower)
			if len(matches) >= 2 {
				version := matches[1]
				fmt.Printf("Version: %s\n", version)
				serviceVersions = append(serviceVersions, service+"/"+version)
			} else {
				serviceVersions = append(serviceVersions, service+"/N")
			}
		}
	}
	//fmt.Println(serviceVersions)
	return serviceVersions
}
