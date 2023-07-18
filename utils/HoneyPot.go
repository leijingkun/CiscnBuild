package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func HoneyPot(ip string, port int) string {
	nc_rev(ip, port)
	//识别kippo(ssh)

	return "hello"

}

func IsKippo(ip string, port int) bool {
	data := []byte("SSH-1.9-OpenSSH_5.9p1\r\n")
	response := nc_send(ip, port, data, 1024)
	// 输出服务器的响应
	fmt.Println("response", response)
	//判断回复里有没有bad字符串
	return strings.Contains(response, "bad")
}
func Isglastopf(ip string, port int) bool {
	response, err := Curl("http://"+ip+":"+strconv.Itoa(port), map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
	if err != nil {
		panic(err)
	}
	fmt.Println("response", response)
	return strings.Contains(response, "Please post your comments for the blog")
}

func IsHFish(ip string, port int) bool {
	response, err := Curl("http://"+ip+":"+strconv.Itoa(port), map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
	if err != nil {
		panic(err)
	}
	fmt.Println("response", response)
	return strings.Contains(response, "/w-logo-blue.png?ver=20131202") && (strings.Contains(response, "?ver=5.2.2"))
}
