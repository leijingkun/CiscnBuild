package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func HoneyPot(ip string) []string {
	result := []string{}
	if IsKippo(ip) != "" {
		result = append(result, IsKippo(ip))
	}
	if Isglastopf(ip) != "" {
		result = append(result, Isglastopf(ip))
	}

	return result
}

// 时间来不及了,只能检测特定端口号
func IsKippo(ip string) string {
	data := []byte("SSH-1.9-OpenSSH_5.9p1\r\n")
	list := []int{22, 2222, 12222, 22222}
	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}
	for _, port := range list {
		wg.Add(1)
		go func(port int) string {
			defer wg.Done()
			response := nc_send(ip, port, data, 1024)
			// 输出服务器的响应
			// fmt.Println("response", response)
			//判断回复里有没有bad字符串
			if strings.Contains(response, "bad") {
				return "kippo" + strconv.Itoa(port)
			}
			return ""
		}(port)
	}
	wg.Wait()
	return ""
}

// glastopf,Hfish都是http协议,放一块写,少发一次curl
func Isglastopf(ip string) string {
	list := []int{80, 8080}
	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}
	for _, port := range list {
		wg.Add(1)
		go func(port int) string {
			defer wg.Done()
			response, err := Curl("http://"+ip+":"+strconv.Itoa(port), map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"})
			if err != nil {
				panic(err)
			}
			fmt.Println("response", response)
			if strings.Contains(response, "Please post your comments for the blog") {
				return "glastopf/" + strconv.Itoa(port)
			}
			if strings.Contains(response, "/w-logo-blue.png?ver=20131202") && (strings.Contains(response, "?ver=5.2.2")) {
				return "hfish/" + strconv.Itoa(port)
			}
			return ""
		}(port)
	}
	wg.Wait()
	return ""
}
