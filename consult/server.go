package consult

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

// 解析客户端的请求
func AnalyticConnection(conn net.Conn) Address {
	readString, err := bufio.NewReader(conn).ReadString('\n')
	var addr Address
	if err != nil {
		fmt.Println("读取请求失败")
		return addr
	}
	err = json.Unmarshal([]byte(readString), &addr)
	if err != nil {
		fmt.Println("解析请求失败")
		return addr
	}
	return addr
}

// 响应代理请求
func Response(conn net.Conn, status Status) error {
	marshal, err := json.Marshal(status)
	if err != nil {
		fmt.Println("将Status转json失败")
		return err
	}
	_, err = conn.Write(append(marshal, '\n'))
	if err != nil {
		fmt.Println("返回数据给客户端失败")
		return err
	}
	return nil
}
