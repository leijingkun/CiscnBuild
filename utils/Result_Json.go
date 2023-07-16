package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONWriter struct {
	f     *os.File
	mutex *sync.Mutex
}

func (jw *JSONWriter) Push(m map[string]interface{}) {
	jw.mutex.Lock()
	defer jw.mutex.Unlock()
	stat, err := jw.f.Stat()
	if err != nil {
		fmt.Println("文件打开错误")
	}
	jsonBuf, _ := json.MarshalIndent(m, "\t", "\t")
	jsonBuf = append(jsonBuf, []byte("\n]\n")...)
	if stat.Size() == 2 {
		jw.f.Seek(stat.Size()-1, 0)
		jsonBuf = append([]byte("\n\t"), jsonBuf...)
		jw.f.Write(jsonBuf)
	} else {
		jw.f.Seek(stat.Size()-4, 0)
		jsonBuf = append([]byte("},\n\t"), jsonBuf...)
		jw.f.Write(jsonBuf)
	}
}
func loadOutputJSON(path string) *JSONWriter {
	if path == "" {
		return nil
	}
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		fmt.Println("检测到JSON输出文件已存在，将自动删除该文件：", path)
		if err := os.Remove(path); err != nil {
			fmt.Println("删除文件失败，请检查：", err)
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR, 0764)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	jw := &JSONWriter{f, &sync.Mutex{}}
	jw.f.Seek(0, 0)
	_, err = jw.f.WriteString(`[]`)
	if err != nil {
		fmt.Println("写入文件失败")
	}
	return &JSONWriter{f, &sync.Mutex{}}
}

func Result(ip string, ports []int) {

	// 定义等待组，用于等待所有协程执行完成
	wg := &sync.WaitGroup{}
	jw := loadOutputJSON("result.json")
	// 使用协程并发识别指定端口的协议
	for _, port := range ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			// 尝试连接指定端口并发送协议识别请求
			server := make(map[string]interface{})
			server["port"] = port
			server["protocol"] = ProtocolDetect(ip, port)
			server["service_app"] = ServiceDetect(ip, port)

			jw.Push(server)

			// 发送协议识别请求，并读取响应数据

		}(port)
	}

	wg.Wait()
}
