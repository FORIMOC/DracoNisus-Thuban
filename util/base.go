package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
)

// GetCallerInfo 获取调用者位置
// 栈帧层级frameLevel => 调用位置字符串
func GetCallerInfo(frameLevel int) string {
	_, file, line, _ := runtime.Caller(frameLevel)
	return fmt.Sprintf("%s, line: %d", file, line)
}

// CombineURL 组合各部分以形成完整URL
// domain, port, uri, query => URL
func CombineURL(domain string, port int, uri string, query string) (string, error) {
	// 将域名和端口号组合为URL字符串
	domainWithPort := fmt.Sprintf("%s:%d", domain, port)

	// 将域名和端口号解析为URL对象
	baseURL, err := url.Parse(domainWithPort)
	if err != nil {
		return "", err
	}

	// 解析URI并组合到基本URL中
	baseURL.Path = uri

	// 解析查询参数并添加到URL中
	queryParams, err := url.ParseQuery(query)
	if err != nil {
		return "", err
	}
	baseURL.RawQuery = queryParams.Encode()

	// 获取完整的URL字符串
	fullURL := baseURL.String()

	return fullURL, nil
}

// SendJSON 发送HTTP POST的JSON请求
// 请求url, 请求json数据 => 请求数据的 body 部分字节流
func SendJSON(url string, jsonData map[string]interface{}) []byte {
	// 编码json
	requestBody, _ := json.Marshal(jsonData)

	// 创建HTTP请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return body
}
