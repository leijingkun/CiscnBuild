# For Ciscn2023 build 
## utils文件夹负责各个功能模块
1. Device.go 负责设备识别,传入参数是ip
2. HoneyPot.go 负责蜜罐识别,传入参数是ip
3. Output_Json.go 负责输出格式化的json文件
4. PortScan 负责扫描端口,打开的端口进行服务和协议识别
5. ServiceDectet.go 负责协议探测
6. Protocol.go 负责协议识别

---
config/config.go负责各项配置
ip列表,top1000端口,协议和服务的范围在里面

---

