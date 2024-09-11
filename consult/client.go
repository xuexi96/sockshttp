package consult

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

// 发送给代理服务器的请求
func Connection(conn net.Conn, addr Address) (Status, error) {
	var status Status

	marshal, err := json.Marshal(addr)
	if err != nil {
		return status, fmt.Errorf("转换Address为json数据失败")
	}

	_, err = conn.Write(append(marshal, '\n'))
	if err != nil {

		return status, fmt.Errorf("发送请求给代理服务器失败")
	}
	readString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return status, fmt.Errorf("读取代理服务器响应数据失败")
	}
	err = json.Unmarshal([]byte(readString), &status)
	if err != nil {
		return status, fmt.Errorf("解析代理服务器响应数据失败")
	}
	return status, nil
}
