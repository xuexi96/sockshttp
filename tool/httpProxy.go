package tool

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func GetHttpAddr(conn net.Conn) string {
	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取CONNECT失败", err)
		return ""
	}
	// 解析 CONNECT 请求
	request = strings.TrimSpace(request)
	parts := strings.Split(request, " ")
	if len(parts) < 3 || parts[0] != "CONNECT" {
		fmt.Println("解析 CONNECT 请求失败")
		return ""
	}
	// 提取目标地址
	targetAddress := parts[1]

	return targetAddress
}

func Answer(conn net.Conn) {
	_, err := fmt.Fprintf(conn, "HTTP/1.1 200 Connection Established\r\n\r\n")
	if err != nil {
		fmt.Println("Error sending response to client:", err)
		return
	}

}
