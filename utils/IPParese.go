package utils

import (
	"CiscnMap/config"
	"bufio"
	"fmt"
	"net"
	"os"
)

func GetScope() {
	file, err := os.Open("ip.txt") // 请替换为你的文件路径
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip, ipnet, err := net.ParseCIDR(scanner.Text())
		if err != nil {
			fmt.Println("Error parsing CIDR: ", err)
			return
		}

		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			config.Hosts = append(config.Hosts, ip.String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	// for _, host := range config.Hosts {
	// 	fmt.Println(host)
	// }
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
