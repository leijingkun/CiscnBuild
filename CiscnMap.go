package main

import "CiscnMap/utils"

func main() {
	//主机存活与否,无论存活加入端口扫描队列
	utils.HostAlive()
	//端口存活与否,存活进入协议识别

	//协议存在与否都进入服务识别

	//设备识别

	//蜜罐识别
	//写上时间戳
}
