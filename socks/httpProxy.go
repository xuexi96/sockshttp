package socks

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func GetHttpProxyAddress(conn net.Conn) (string, error) {
	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("读取CONNECT失败")
	}
	// 解析 CONNECT 请求
	request1 := request
	request = strings.TrimSpace(request)
	parts := strings.Split(request, " ")
	if len(parts) < 3 || parts[0] != "CONNECT" {
		fmt.Println(request1)
		return "", fmt.Errorf("解析 CONNECT 请求失败")
	}
	// 提取目标地址
	targetAddress := parts[1]
	return targetAddress, nil
}

func ResponseHttpProxy(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "HTTP/1.1 200 Connection Established\r\n\r\n")
	if err != nil {
		fmt.Println("Error sending response to client:", err)
		return err
	}
	return nil

}
