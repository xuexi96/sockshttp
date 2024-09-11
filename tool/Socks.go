package tool

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type targetaddress struct {
	address string
	port    uint16
}

func Socksnegotiate(conn net.Conn) {

}

// socks协商
func handleSocks5Handshake(conn net.Conn) error {
	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("读取协商信息失败: %w", err)
	}

	// 检查socks版本
	if buf[0] != 0x05 {
		return fmt.Errorf("检查socks版本不是socks5", err)
	}

	//响应客户端
	_, err = conn.Write([]byte{0x05, 0x00})
	if err != nil {
		return fmt.Errorf("响应客户端失败:%w", err)
	}

	log.Printf("SOCKS5 握手成功 (客户端支持 %d 种认证方法)", buf[1])
	return nil

}

// 请求连接
func getAddr(conn net.Conn) *targetaddress {
	buf := make([]byte, 256)
	var target targetaddress
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取请求失败")
		return nil
	}

	if buf[0] != 0x05 {
		fmt.Println("请求不是socks5")
		return nil
	}

	// 检查请求类型（0x01 表示 CONNECT）
	if buf[1] == 0x01 {
		fmt.Println("请求类型为CONNECT")
	} else if buf[1] == 0x02 {
		fmt.Println("请求类型为BIND,当前程序未实现该功能")
		return nil
	} else if buf[1] == 0x03 {
		fmt.Println("请求类型为UDP ASSOCIATE,当前程序未实现该功能")
		return nil
	} else {
		fmt.Println("未失的请求")
		return nil
	}

	// 判断目标地址类型
	// 读取目标地址和端口
	addrType := buf[3]
	switch addrType {
	case 0x01: // IPv4
		target.address = net.IP(buf[4:8]).String()
	case 0x03: // 域名
		target.address = string(buf[5 : 5+buf[4]])
	case 0x04: // IPv6
		target.address = net.IP(buf[4:20]).String()
	default:
		fmt.Println("不支持的地址类型: %d", addrType)
		return nil
	}
	target.port = binary.BigEndian.Uint16(buf[n-2 : n])
	log.Println("获取目标地址及端口号成功")
	return &target
}

func Successfulresponse(conn net.Conn) {

	conn.Write([]byte{0x05, 0x00})
}
