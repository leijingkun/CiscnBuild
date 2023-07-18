package utils

import (
	"strconv"
	"strings"

	"github.com/lcvvvv/gonmap"
)

// "ssh", "http", "https", "rtsp", "ftp", "telnet", "amap", "mongodb", "redis", "mysql"

// 根据端口和IP确定协议
func ProtocolDetect(ip string, port int) string {
	return gonmap.GuessProtocol(port)
}
func GuessProtocol(port int) string {
	protocol := nmapServices[port]
	if protocol == "unknown" {
		protocol = "http"
	}
	return protocol
}

var nmapServicesString = `20	ftp
21	ftp
22	ssh
23	telnet
25	smtp
80	http
443	https
5671	amqp
5672	amqp
22	ssh
322	rtsp
554	rtsp
8554	rtsp
21	ftp
26	ftp
47	ftp
69	ftp
63	ftp
115	ftp
152	ftp
349	ftp
574	ftp
990	ftps
992	telnet
6623	telnet
6379	redis
3306	mysql
2273	mysql
6446	mysql
33060	mysql`

var nmapServices = func() []string {
	var r []string
	for _, line := range strings.Split(nmapServicesString, "\n") {
		index := strings.Index(line, "\t")
		v1 := line[:index]
		v2 := line[index+1:]
		port, _ := strconv.Atoi(v1)
		protocol := v2

		for i := len(r); i < port; i++ {
			r = append(r, "unknown")
		}

		protocol = FixProtocol(protocol)
		r = append(r, protocol)
	}
	return r
}()

func FixProtocol(oldProtocol string) string {
	//进行最后输出修饰
	if oldProtocol == "ssl/http" {
		return "https"
	}
	if oldProtocol == "http-proxy" {
		return "http"
	}
	if oldProtocol == "ms-wbt-server" {
		return "rdp"
	}
	if oldProtocol == "microsoft-ds" {
		return "smb"
	}
	if oldProtocol == "netbios-ssn" {
		return "netbios"
	}
	if oldProtocol == "oracle-tns" {
		return "oracle"
	}
	if oldProtocol == "msrpc" {
		return "rpc"
	}
	if oldProtocol == "ms-sql-s" {
		return "mssql"
	}
	if oldProtocol == "domain" {
		return "dns"
	}
	if oldProtocol == "svnserve" {
		return "svn"
	}
	if oldProtocol == "ibm-db2" {
		return "db2"
	}
	if oldProtocol == "socks-proxy" {
		return "socks5"
	}
	if len(oldProtocol) > 4 {
		if oldProtocol[:4] == "ssl/" {
			return oldProtocol[4:] + "-ssl"
		}
	}
	oldProtocol = strings.ReplaceAll(oldProtocol, "_", "-")
	return oldProtocol
}
