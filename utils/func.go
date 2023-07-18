package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 返回数组里随机一个元素
func randomFrom(arr []string) string {
	n := len(arr)

	// 生成0到n-1之间的随机数字
	idx := rand.Intn(n)

	return arr[idx]
}

// 模拟nc,发送数据
func nc_send(ip string, port int, data []byte, buf int) string {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("连接失败：", err)
		return "error"
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送数据失败：", err)
		return "send error"
	}
	// 接收服务器的响应
	buffer := make([]byte, buf)
	_, err = conn.Read(buffer)
	if err != nil {
		// 处理读取错误
		return "error"
	}

	// 打印返回数据
	response := string(buffer)
	fmt.Println("Received", response)
	return response
}

// 接收数据
func nc_rev(ip string, port int) string {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(""))
	defer conn.Close() // Close the connection when we're done

	buf := make([]byte, 1024)
	//等待结果,超时3s则停止
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Set a 5-second timeout
		n, err := conn.Read(buf)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				break // Exit loop when we reach the timeout
			}
			if err != io.EOF {
				panic(err)
			}
			break // Exit loop when we reach end of file
		}
		fmt.Println(string(buf[:n]))
	}
	return ""
}

// 模拟curl
func Curl(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// 设置请求头部信息
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将响应的内容转换为string类型并返回
	return string(body), nil
}
