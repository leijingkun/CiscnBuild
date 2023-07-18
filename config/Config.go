package config

var (
	//主机范围
	Hosts = []string{}
	//探测的范围
	Server_list   = []string{"Windows", "centos", "ubuntu", "openssh", "openssl", "wordpress", "LiteSpeed", "Jetty", "java", "node.js", "express", "asp.net", "php", "Microsoft-HTTPAPI", "rabbitmp", "apache", "iis", "nginx", "micro_httpd", "openresty", "grafana", "Weblogic", "elasticsearch", "debian"}
	Honeypot_list = map[string]string{
		"webcam":   "Hikvision",
		"firewall": "pfsense",
		"switch":   "cisco",
		"nas":      "synology",
	}
	Device_list   = []string{"pfsense", "Hikvision", "cisco", "dahua", "synology"}
	Protocol_list = []string{"ssh", "http", "https", "rtsp", "ftp", "telnet", "amap", "mongodb", "redis", "mysql"}
	//常用端口列表
	Ports = []int{7, 9, 13, 21, 22, 25, 37, 53, 79, 80, 88, 106, 110, 113, 119, 135, 139, 143, 179, 199, 389, 427, 443,
		444,
		445, 465, 513, 514, 543, 548, 554, 587, 631, 646, 873, 990, 993, 995, 1025, 1026, 1027, 1028, 1110,
		1433,
		1720, 1723, 1755, 1900, 2000, 2049, 2121, 2717, 3000, 3128, 3306, 3389, 3986, 4899, 5000, 5009, 5051,
		5060, 5101, 5190, 5357, 5432, 5631, 5666, 5800, 5900, 6000, 6646, 7070, 8000, 8008, 8080, 8443, 8888,
		9100, 9999, 32768, 49152, 49153, 49154, 49155, 49156}
)
