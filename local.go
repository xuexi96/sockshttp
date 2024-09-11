package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"sockshttp/consult"
	"sockshttp/socks"
	"sync"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		listen, err := listen.Accept()
		if err != nil {
			fmt.Println("监听数据失败：", err)
		}
		go Proxy(listen)

	}

}

func Proxy(conn net.Conn) {
	defer conn.Close()
	address, err := socks.GetHttpProxyAddress(conn)
	if err != nil {
		fmt.Println("获取代理失败", err)
		return
	}
	dial, err := tls.Dial("tcp", "103.38.83.233:8081", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("已代理服务器建立tls服务器失败")
		return
	}
	defer dial.Close()
	var addr consult.Address
	addr.RemoteAddr = address
	addr.Remoteprotocoltype = "tcp"
	status, err := consult.Connection(dial, addr)
	if err != nil {
		fmt.Println("连接到目标服务器失败：", err)
		return
	}
	err = socks.ResponseHttpProxy(conn)
	if err != nil {

		fmt.Println("响应http代理失败")
		return
	}
	fmt.Println(dial.LocalAddr().String() + "<<------------------------------------>>" + status.LocalAddr)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := io.Copy(dial, conn) // 客户端 -> 目标服务器
		if err != nil {
			fmt.Println("从客户端到目标服务器的数据传输失败:", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = io.Copy(conn, dial) // 目标服务器 -> 客户端
		if err != nil {
			fmt.Println("从目标服务器到客户端的数据传输失败:", err)
			return
		}
	}()
	wg.Wait()

}
