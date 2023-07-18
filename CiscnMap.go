package main

import (
	"CiscnMap/utils"
)

func main() {
	//读取ip.txt获取范围
	utils.GetScope()
	//测绘,启动!
	utils.HostAlive()
}
