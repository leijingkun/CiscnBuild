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

// 负责写入文件,是追加格式,
func (jw *JSONWriter) Push(m map[string]interface{}) {
	jw.mutex.Lock()
	defer jw.mutex.Unlock()
	stat, err := jw.f.Stat()
	if err != nil {
		fmt.Println("打开文件失败")
	}
	newmap := make(map[string]interface{})
	newmap[m["ip"].(string)] = map[string]interface{}{
		"services":   m["services"],
		"deviceinfo": m["deviceinfo"],
		"honeypot":   m["honeypot"],
	}
	jsonBuf, _ := json.MarshalIndent(newmap, "\t", "\t")
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

// 打开一个json文件,如果存在也不会删除,因为需要多次写入同一个文件
func loadOutputJSON(path string) *JSONWriter {
	if path == "" {
		return nil
	}
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		fmt.Println("检测到JSON输出文件已存在，将自动删除该文件：", path)
		//删除文件,先注释掉

		// if err := os.Remove(path); err != nil {
		// 	fmt.Println("删除文件失败，请检查：", err)
		// }
	}
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR, 0764)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	jw := &JSONWriter{f, &sync.Mutex{}}
	jw.f.Seek(0, 0)
	_, err = jw.f.WriteString(``)
	if err != nil {
		fmt.Println("写入文件失败")
	}
	return &JSONWriter{f, &sync.Mutex{}}
}

func Result(ip string, port int, ipInfo map[string]interface{}) {
	if ipInfo["services"] == nil {
		ipInfo["services"] = make([]map[string]interface{}, 0)
	}
	//设置一个services的三个字段
	ipInfo["services"] = append(ipInfo["services"].([]map[string]interface{}), map[string]interface{}{
		"port":        port,
		"protocol":    ProtocolDetect(ip, port),
		"service_app": ServiceDetect(ip, port),
	})

}
